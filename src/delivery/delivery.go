package delivery

type (
	Shippable interface {
		Load() *Shippable
		Ship()
	}

	ShippableCargo interface {
		Load()
	}
)
