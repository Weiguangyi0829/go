package main

func main() {
	{
		a := 1
		if false {

		} else {
			{
				b := 2
				if false {

				} else {

					{
						c := 3
						if false {

						} else {
							println(a, b, c)
						}
					}
				}
			}
		}
	}
}
