resource "cloudflare_zone" "dns_zone" {
  zone = "wspieramprzyrode.pl"
  plan = "free"
}