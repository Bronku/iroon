# TODO List for Iroon Project

## Critical Features

- [x] Database Implementation

  - [x] Init mode to create a new db
  - [x] Set up SQLite database
  - [x] Create proper schemas
  - [x] Implement data persistence
  - [x] Add database migrations
  - [x] Add database caching

- [ ] Input Validation

  - [ ] Phone number format validation
  - [ ] Required fields validation
  - [ ] Date validation (no past dates)
  - [ ] Cake amount limits
  - [ ] Form validation feedback

- [ ] Order Form
  - [ ] table headers
  - [ ] padding
  - [ ] Prevent enter from submitting the form
  - [ ] fix basket scrolling

- [ ] Price Management
  - [x] Calculate order totals
  - [x] Handle partial payments
  - [ ] Add payment status tracking
  - [ ] Implement payment validation

## Security

- [x] Authentication System
  - [x] Persistent logins across server restarts
  - [x] User login
  - [x] User logout
  - [ ] Role-based access (admin, staff)
  - [x] Password hashing
- [ ] CSRF Protection
- [ ] Input Sanitization
- [x] Session Management

## UI/UX Improvements

- [ ] Add CSS Styling
  - [ ] Responsive design
  - [ ] Dark/Light theme
  - [ ] Print-friendly styles
- [ ] Improve Navigation
- [ ] Add Loading States
- [x] Confirmation Dialogs
- [ ] Form Autosave
- [x] Success/Error Notifications

## Order Management

- [x] Order Status System
- [ ] Order Search & Filtering
  - [ ] By date range
  - [ ] By status
  - [ ] By customer
- [ ] Order History
- [ ] Order Notes/Comments
- [ ] Bulk Order Operations

## Inventory Management

- [ ] Track Ingredients
- [ ] Low Stock Alerts
- [x] Cake Catalog Management
  - [x] Add/Edit/Delete cakes
  - [x] Cake categories
  - [ ] Multiple categories for each cake
  - [ ] Cake images
- [x] Seasonal Menu Items

## Reporting & Analytics

- [ ] Sales Reports
- [ ] Popular Items
- [ ] Customer Analytics
- [ ] Revenue Reports
- [ ] Export Data (CSV/PDF)

## Customer Features

- [ ] Customer Accounts
- [ ] Order History
- [ ] Favorite Items
- [ ] Loyalty Program
- [ ] Reviews & Ratings

## Communication

- [ ] Email Notifications
  - [ ] Order confirmation
  - [ ] Status updates
  - [ ] Payment reminders
- [ ] SMS Notifications
- [ ] Print Order Slips

## Additional Features

- [ ] Calendar View
  - [ ] Daily orders
  - [ ] Capacity planning
- [ ] Delivery Management
  - [ ] Delivery zones
  - [ ] Delivery fees
  - [ ] Driver assignments
- [ ] Special Requirements
  - [ ] Allergies
  - [ ] Custom decorations
  - [ ] Special messages
- [ ] Multi-language Support
- [ ] Backup System

## Technical Improvements

- [ ] Error Handling
  - [ ] Proper error logging
  - [ ] User-friendly error messages
  - [ ] A better way to handle the error page
- [ ] Performance Optimization
- [ ] API Documentation
- [ ] Unit Tests
- [ ] Integration Tests
- [x] File structure
- [x] Embed templates
- [x] Single init function for handler

## Future Considerations

- [ ] Mobile App
- [ ] Online Payment Integration
- [ ] Integration with POS Systems
- [ ] Customer Feedback System
- [ ] Marketing Tools
  - [ ] Promotions
  - [ ] Discount codes
  - [ ] Gift cards
