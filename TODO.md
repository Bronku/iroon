# TODO List for Iroon Project

## Critical Features

- [ ] Database Implementation

  - [ ] Init mode to create a new db
  - [x] Set up SQLite database
  - [x] Create proper schemas
  - [x] Implement data persistence
  - [ ] Add database migrations
  - [ ] Add database caching

- [ ] Input Validation

  - [ ] Phone number format validation
  - [ ] Required fields validation
  - [ ] Date validation (no past dates)
  - [ ] Cake amount limits
  - [ ] Form validation feedback

- [ ] Price Management
  - [ ] Calculate order totals
  - [ ] Handle partial payments
  - [ ] Add payment status tracking
  - [ ] Implement payment validation

## Security

- [ ] Authentication System
  - [ ] User login/logout
  - [ ] Role-based access (admin, staff)
  - [ ] Password hashing
- [ ] CSRF Protection
- [ ] Input Sanitization
- [ ] Session Management

## UI/UX Improvements

- [ ] Add CSS Styling
  - [ ] Responsive design
  - [ ] Dark/Light theme
  - [ ] Print-friendly styles
- [ ] Improve Navigation
- [ ] Add Loading States
- [ ] Confirmation Dialogs
- [ ] Form Autosave
- [ ] Success/Error Notifications

## Order Management

- [ ] Order Status System
  - [ ] Pending
  - [ ] In Progress
  - [ ] Completed
  - [ ] Cancelled
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
- [ ] Cake Catalog Management
  - [ ] Add/Edit/Delete cakes
  - [ ] Cake categories
  - [ ] Cake images
- [ ] Seasonal Menu Items

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
- [ ] Performance Optimization
- [ ] API Documentation
- [ ] Unit Tests
- [ ] Integration Tests
- [ ] File structure
- [ ] Embed templates
- [ ] Single init function for handler

## Future Considerations

- [ ] Mobile App
- [ ] Online Payment Integration
- [ ] Integration with POS Systems
- [ ] Customer Feedback System
- [ ] Marketing Tools
  - [ ] Promotions
  - [ ] Discount codes
  - [ ] Gift cards
