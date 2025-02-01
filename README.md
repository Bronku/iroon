# iroon
## todo:
- [x] extract order creation form into it's own element
- [ ] organize html structure
- [x] option to add additional cakes
- [x] saving orders on server
- [ ] displaying orders
- [x] move id to cake struct
- [ ] refresh the list on submission
- [ ] peristent storage
- [x] rudimentary css
- [ ] organize the cake editor
## ux design
### adding an order:
1. user selects the _add order_ button
    1. the button should open an order creation menu
2. user enters the customer details
    1. incorrect details are highlighted
    2. user can't submit an incorrect order
3. user selects the cakes to add from the list
    1. the list should be scrollable, with a search box if needed
    2. added cakes should appear next to other order details
    3. the user can change the count, or remove the cake entirely
4. user submits the order
    1. after submission the user should get a confirmation, and an order id along with the link to said order page
    2. the form should disappear after submission
    3. there should be a button to go back to creating a new form
### adding a new cake: 
1. user selects the _new cake_ button
2. enters the new details
3. user selects submit
    