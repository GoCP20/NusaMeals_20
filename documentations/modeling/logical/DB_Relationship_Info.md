# DB Relationship Info

- Tabel User: 1-to-1 dengan Tabel Level (1 User memiliki 1 Level).

- Tabel Level: Tidak ada relasi dengan tabel lain.

- Tabel Menu: 1-to-1 dengan Tabel Category (1 Menu terkait dengan 1 Category).

- Tabel Category: Tidak ada relasi dengan tabel lain.

- Tabel Food: 1-to-1 dengan Tabel Category (1 Food terkait dengan 1 Category).

- Tabel Drink: 1-to-1 dengan Tabel Category (1 Drink terkait dengan 1 Category).

- Order: 1-to-N dengan Tabel Customer (1 Order dimiliki oleh 1 Customer, 1 Customer dapat memiliki banyak Order), 1-to-1 dengan Tabel Menu (1 Order terkait dengan 1 Menu), 1-to-1 dengan Tabel Payment_Type (1 Order terkait dengan 1 Payment_Type).

- Tabel Order_Cancel: 1-to-1 dengan Tabel Order (1 Order_Cancel terkait dengan 1 Order), 1-to-1 dengan Tabel Customer (1 Order_Cancel terkait dengan 1 Customer).

- Tabel Reservation: Tidak ada relasi dengan tabel lain.

- Tabel List_Table: Tidak ada relasi dengan tabel lain.

- Tabel Report: Tidak ada relasi dengan tabel lain.

- Tabel Customer: 1-to-1 dengan Tabel User (1 Customer terkait dengan 1 User), 1-to-N dengan Tabel Order (1 Customer memiliki banyak Order).

- Tabel Customer_Ewallet_Account: 1-to-1 dengan Tabel Ewallet_Type (1 Customer_Ewallet_Account terkait dengan 1 Ewallet_Type), 1-to-1 dengan Tabel Customer (1 Customer_Ewallet_Account terkait dengan 1 Customer).

- Tabel Customer_Card_Account: 1-to-1 dengan Tabel Card_Type (1 Customer_Card_Account terkait dengan 1 Card_Type), 1-to-1 dengan Tabel Customer (1 Customer_Card_Account terkait dengan 1 Customer).

- Tabel Payment: 1-to-1 dengan Tabel Order (1 Payment terkait dengan 1 Order), 1-to-1 dengan Tabel Payment_Type (1 Payment terkait dengan 1 Payment_Type).

- Tabel Payment_Type: Tidak ada relasi dengan tabel lain.

- Tabel Card_Type: Tidak ada relasi dengan tabel lain.

- Tabel Ewallet_Type: Tidak ada relasi dengan tabel lain.
