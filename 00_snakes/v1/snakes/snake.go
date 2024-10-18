package snakes

type SnakeBody struct {
	X      int
	Y      int
	Xspeed int
	Yspeed int
}

func (sb *SnakeBody) ChangeDir(horizontal int, vertical int) {
	sb.Xspeed = horizontal
	sb.Yspeed = vertical
}

func (sb *SnakeBody) Update(width int, height int) {
	sb.X = (sb.X + sb.Xspeed) % width
	if sb.X < 0 {
		sb.X += width
	}

	sb.Y = (sb.Y + sb.Yspeed) % height
	if sb.Y < 0 {
		sb.Y += height
	}
}
