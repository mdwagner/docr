package crystal

func FindType(lib *Type, symbol string) any {
	types := []*Type{lib}
	if len(lib.Types) != 0 {
		types = append(types, lib.Types...)
	}

	return findType(types, symbol)
}

func findType(types []*Type, symbol string) any {
	var overflow []*Type

	for _, t := range types {
		if t.Name == symbol || t.FullName == symbol {
			return t
		}

		for _, c := range t.Constants {
			if c.Name == symbol {
				return c
			}
		}

		for _, c := range t.Constructors {
			if c.Name == symbol {
				return c
			}
		}

		for _, m := range t.ClassMethods {
			if m.Name == symbol {
				return m
			}
		}

		for _, m := range t.InstanceMethods {
			if m.Name == symbol {
				return m
			}
		}

		for _, m := range t.Macros {
			if m.Name == symbol {
				return m
			}
		}

		overflow = append(overflow, t.Types...)
	}

	if len(overflow) == 0 {
		return nil
	}

	return findType(overflow, symbol)
}
