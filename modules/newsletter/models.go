package newsletter

type Subscription struct {
	UUID string `gorm:"primary_key" json:"uuid"`
}
