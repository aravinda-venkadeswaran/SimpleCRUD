package models

type User struct {
	ID        uint    `json:"id" gorm:"primarykey"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Address   Address `json:"address" gorm:"foreignkey:AddressID"`
	Phone     string  `json:"phone"`
	Website   string  `json:"website"`
	Company   Company `json:"company" gorm:"foreignkey:CompanyID"`
	AddressID uint    `json:"-"`
	CompanyID uint    `json:"-"`
	Posts     []Post  `gorm:"foreignkey:UserID"`
}

type Address struct {
	ID      uint   `json:"-" gorm:"primarykey"`
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo" gorm:"foreignkey:GeoID"`
	GeoID   uint   `json:"-"`
}

type Company struct {
	ID          uint   `json:"-" gorm:"primarykey"`
	Name        string `json:"name"`
	CatchPhrase string `json:"catchphrase"`
	Bs          string `json:"bs"`
}

type Geo struct {
	ID  uint   `json:"-" gorm:"primarykey"`
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Post struct {
	UserID   uint      `json:"userId"`
	ID       uint      `json:"id" gorm:"primarykey"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Comments []Comment `gorm:"foreignkey:PostID"`
}

type Comment struct {
	PostID uint   `json:"postId"`
	ID     uint   `json:"id" gorm:"primarykey"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
