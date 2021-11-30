package twitch

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/nicklaw5/helix"
)

func resourceChannelPoint() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceChannelPointCreate,
		ReadContext:   resourceChannelPointRead,
		DeleteContext: resourceChannelPointDelete,

		Schema: map[string]*schema.Schema{
			"broadcaster_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The broadcaster you want to create this channel point for",
				Required:    true,
				ForceNew:    true,
			},
			"title": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The title of the channel point",
				Required:    true,
				ForceNew:    true,
			},
			"cost": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "The cost of the channel point",
				Required:    true,
				ForceNew:    true,
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Is this prompt redeemable",
				Default:     true,
				ForceNew:    true,
				Optional:    true,
			},
			"prompt": &schema.Schema{
				Type:        schema.TypeString,
				Description: "The prompt for the viewer when redeeming the reward.",
				Optional:    true,
				ForceNew:    true,
			},
			"max_per_stream": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "how many instances of this reward can redeemed within a stream",
				Optional:    true,
				ForceNew:    true,
			},
			"global_cooldown": &schema.Schema{
				Type:        schema.TypeInt,
				Description: "Whether a cooldown is enabled and what the cooldown is.",
				Optional:    true,
				ForceNew:    true,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceChannelPointCreate(ctx context.Context, d *schema.ResourceData, c interface{}) diag.Diagnostics {
	client := c.(*helix.Client)
	req := &helix.ChannelCustomRewardsParams{
		BroadcasterID: d.Get("broadcaster_id").(string),
		Title:         d.Get("title").(string),
		Cost:          d.Get("cost").(int),
		IsEnabled:     d.Get("enabled").(bool),
	}

	if v, b := d.GetOk("prompt"); b {
		req.Prompt = v.(string)
	}

	if v, b := d.GetOk("max_per_stream"); b {
		req.MaxPerStream = v.(int)
		req.IsMaxPerStreamEnabled = true
	}

	if v, b := d.GetOk("global_cooldown"); b {
		req.GlobalCooldownSeconds = v.(int)
		req.IsGlobalCooldownEnabled = true
	}

	// rewardResponse, err := client.CreateCustomReward(&helix.ChannelCustomRewardsParams{
	// 	BroadcasterID: "720264417",
	// 	Title:         "hi chat",
	// 	Cost:          500000,
	// 	IsEnabled:     true,
	// })
	rewardResponse, err := client.CreateCustomReward(req)

	if err != nil {
		// handle error
	}

	fmt.Printf("%+v\n", rewardResponse)

	id := rewardResponse.Data.ChannelCustomRewards[0].ID
	idReal := d.Get("broadcaster_id").(string) + "@" + id
	d.SetId(idReal)
	resourceChannelPointRead(ctx, d, c)
	return nil
}

func resourceChannelPointRead(ctx context.Context, d *schema.ResourceData, c interface{}) diag.Diagnostics {
	client := c.(*helix.Client)

	realId := d.Id()
	ar := strings.Split(realId, "@")

	broadcasterId := ar[0]
	id := ar[1]

	resp, err := client.GetCustomRewards(&helix.GetCustomRewardsParams{
		BroadcasterID: broadcasterId,
		ID:            id,
	})

	if err != nil {
		d.SetId("")
	}

	fmt.Printf("%v\n", resp)

	reward := resp.Data.ChannelCustomRewards[0]

	if err := d.Set("broadcaster_id", reward.BroadcasterID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("title", reward.Title); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("cost", reward.Cost); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceChannelPointDelete(ctx context.Context, d *schema.ResourceData, c interface{}) diag.Diagnostics {
	client := c.(*helix.Client)

	realId := d.Id()
	ar := strings.Split(realId, "@")

	broadcasterId := ar[0]
	id := ar[1]

	_, err := client.DeleteCustomRewards(&helix.DeleteCustomRewardsParams{
		BroadcasterID: broadcasterId,
		ID:            id,
	})
	if err != nil {
		// handle error
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return nil
}
