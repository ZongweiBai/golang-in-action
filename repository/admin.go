package repository

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// User的构造函数,函数不属于任何类型
func NewUser(id uint64, name string) *User {
	return &User{
		ID: id,
		Name: name,
	}
}

// User改变ID的方法,方法属于特定的类型
// 值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身
func (u User) ChangeId(id uint64, user *User) {
	user.ID = id
}

// User改变Name的方法
// 什么时候应该使用指针类型接收者
// 需要修改接收者中的值
// 接收者是拷贝代价比较大的大对象
// 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者
func (u *User) ChangeName(name string) {
	u.Name = name
}

