package model

type User struct {
	Name     string `gorm:"type:varchar(255);not null"`
	LastName string `gorm:"type:varchar(255);not null"`
	UserName string `gorm:"type:varchar(255);not null;unique"`
	Email    string `gorm:"type:varchar(255);not null"`
	Pwd      string `gorm:"type:varchar(255);not null"`
}

type Users []User

/*`id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8_bin NOT NULL,
  `lastName` varchar(255) COLLATE utf8_bin NOT NULL,
  `username`	varchar(255) COLLATE utf8_bin NOT NULL unique,
  `pwd` varchar(255) COLLATE utf8_bin NOT NULL,
  `mail` varchar(255) COLLATE utf8_bin NOT NULL*/
