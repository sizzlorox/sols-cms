# Sols CMS

TODO
* Add user
* Add RBAC
  * Needs to handle a superadmin
* Add Auth Middleware
  * Superadmin needs to override everything
* Add Content Types
* Add Fields
* Add Assets
  * Needs to handle compression
  * Needs to validate asset
* Add Content
  * Needs to support drafts and version control

Admin
POST   /api/domains                          (superadmin)
GET    /api/domains                          (list mine)
GET    /api/domains/:domainSlug              (owner/admin)
PUT    /api/domains/:domainSlug              (owner/admin)
DELETE /api/domains/:domainSlug              (superadmin)

Modelling
POST   /api/:domainSlug/content-types
GET    /api/:domainSlug/content-types
GET    /api/:domainSlug/content-types/:ctSlug
PUT    /api/:domainSlug/content-types/:ctSlug
DELETE /api/:domainSlug/content-types/:ctSlug

POST   /api/:domainSlug/content-types/:ctSlug/fields
PUT    /api/:domainSlug/fields/:fieldId
PATCH  /api/:domainSlug/fields/:fieldId/order
DELETE /api/:domainSlug/fields/:fieldId

Content
POST   /api/:domainSlug/content/:ctSlug
GET    /api/:domainSlug/content/:ctSlug
GET    /api/:domainSlug/content/:ctSlug/:id
PUT    /api/:domainSlug/content/:ctSlug/:id
POST   /api/:domainSlug/content/:ctSlug/:id/publish

Assets
POST   /api/:domainSlug/uploads
GET    /api/:domainSlug/assets

Webhooks
POST   /api/:domainSlug/webhooks
