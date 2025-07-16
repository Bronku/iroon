# TODO List for Iroon Project

## critical
- [ ] HTTPS
- [ ] Input Validation
- [ ] Better error Handling
- [ ] CSRF Protection

## architecture
- [ ] Saparate routing logic, and http handlers
- [ ] Proper session auth ?
- [ ] defining protected, and public routes
- [ ] permission System
- [ ] configurable session expiration
- [ ] Storing default login in config file/env variable
- [ ] rework server to simplify loading templates, and routes
- [ ] make fatcher only return data, and error
- [ ] maybe sending forms as json to simplify code?
- [ ] move basket_element_template to be together with script
- [ ] make fetchers only return data, and error, and deduce the http response based on that

## new features
- [ ] Backup System
- [ ] daily summary page
- [ ] filter out done orders
- [ ] Order Notes/Comments
- [ ] Bulk Order Operations
- [ ] Track Ingredients
- [ ] Print Order Slips
- [ ] Capacity planning
- [ ] Delivery Management
- [ ] Custom decorations

## cosmetic
- [ ] /orders/search htmx trigger `<form class="search-container" hx-get="/orders/search" ...>`
- [ ] orders.gohtml remove inline css
- [ ] move style for order_fts into a separate css file
- [ ] move new cake/order into cake/order pages
- [ ] new cake header instead of cake 0
- [ ] form css
- [ ] proper confirmation pages (maybe popus with htmx?)
- [ ] Form validation feedback
- [ ] Prevent enter from submitting the form
- [ ] Mobile interface

## nice to have
- [ ] ci/cd pipeline (github actions)
- [ ] monitoring (prometheus, grafana)
- [ ] structured logs (JSON format), log formatter
