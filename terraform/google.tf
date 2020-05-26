resource "google_firebase_project" "firebase_project" {
  provider = google-beta
}
resource "google_firebase_project_location" "firebase_location" {
  provider = google-beta
  project  = google_firebase_project.firebase_project.project

  location_id = "europe-west"
}
resource "google_firebase_web_app" "firebase_app" {
  provider     = google-beta
  display_name = "WspierajPrzyrode"

  depends_on = [google_firebase_project.firebase_project, google_firebase_project_location.firebase_location]
}

data "google_firebase_web_app_config" "firebase_app" {
  provider   = google-beta
  web_app_id = google_firebase_web_app.firebase_app.app_id
}
