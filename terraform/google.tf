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
  display_name = "WspieramPrzyrode"

  depends_on = [google_firebase_project.firebase_project, google_firebase_project_location.firebase_location]
}

data "google_firebase_web_app_config" "firebase_app" {
  provider   = google-beta
  web_app_id = google_firebase_web_app.firebase_app.app_id
}

resource "google_service_account" "wpieramprzyrode_api_dev_service_account" {
  provider     = google-beta
  account_id   = "wpieramprzyrode-api-dev"
  display_name = "Wspieram Przyrode API dev account"
}

resource "google_service_account_key" "wpieramprzyrode_api_dev_key" {
  provider           = google-beta
  service_account_id = google_service_account.wpieramprzyrode_api_dev_service_account.name
}

resource "google_project_iam_binding" "wspieramprzyrode_api_dev_role_datastore_binding" {
  provider = google-beta
  role     = "roles/datastore.user"
  members = [
         "serviceAccount:${google_service_account.wpieramprzyrode_api_dev_service_account.email}"
  ]
}
