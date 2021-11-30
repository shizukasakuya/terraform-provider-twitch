terraform {
  required_providers {
    twitch = {
      source  = "shizukasakuya/twitch"
    }
  }
} 

provider "twitch" {
  
}

resource "twitch_channel_point" "stretch" {
  broadcaster_id   = "720264417"
  title = "stretch"
  global_cooldown = 900
  cost = 300
}

resource "twitch_channel_point" "hydrate" {
  broadcaster_id   = "720264417"
  title = "hydrate"
  global_cooldown = 900
  cost = 300
}

resource "twitch_channel_point" "daily_kuya_tax" {
  broadcaster_id   = "720264417"
  title = "kuya tax"
  prompt = "does absolutely nothing"
  cost = 1000
}


