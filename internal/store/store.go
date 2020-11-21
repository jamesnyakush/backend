package store

type Store interface {
	Users() UserStore
	Bill() BillStore
	Building() BuildingStore
	County() CountyStore
	Feedback() FeedbackStore
	House() HouseStore
	Mpesa() MpesaStore
	Notice() NoticeStore
	Payment() PaymentStore
	Permission() PermissionStore
	PromoCode() PromoCodeStore
	Role() RoleStore
	Utility() UtilityStore
}

type UserStore interface {
	New()
	Get()
}

type BillStore interface {
	Create()
}

type BuildingStore interface {
}

type CountyStore interface {
}

type FeedbackStore interface {
}

type HouseStore interface {
}

type MpesaStore interface {
}

type NoticeStore interface {
}

type PaymentStore interface {
}

type PermissionStore interface {
}

type PromoCodeStore interface {
}

type RoleStore interface {
}

type UtilityStore interface {
}
