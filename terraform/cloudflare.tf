resource "cloudflare_zone" "dns_zone" {
  zone = "wspieramprzyrode.pl"
  plan = "free"
}

resource "cloudflare_record" "docs_dns_record" {
  zone_id = cloudflare_zone.dns_zone.id
  name    = "docs"
  value   = "wspieramprzyrode.github.io"
  type    = "CNAME"
  ttl     = 3600
}