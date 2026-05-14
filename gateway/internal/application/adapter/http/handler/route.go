package handler

func (c *ClassroomHandler) MapClassroomRoutes() {
	c.group.POST("", c.Create)
	c.group.GET("", c.Read)
	c.group.PUT("", c.Update)
	c.group.DELETE("/:roomNumber", c.Delete)
}

func (l *LockerHandler) MapLockerRoutes() {
	l.group.POST("", l.Create)
	l.group.GET("/:number", l.Read)
	l.group.PUT("", l.Update)
	l.group.DELETE(":number", l.Delete)
}

func (t *TeacherHandler) MapTeacherRoutes() {
	t.group.POST("", t.Create)
	t.group.GET("/:name", t.Create)
	t.group.PUT("", t.Update)
	t.group.DELETE("/:name", t.Delete)
}
