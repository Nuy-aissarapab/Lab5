package entity

import (
	"testing"

	."github.com/onsi/gomega"
)

// 1) กรณีข้อมูลครบทุก field -> ต้องผ่าน
func TestStudent_AllFieldsValid(t *testing.T) {
	g := NewGomegaWithT(t)

	student := Student{
		Fullname: "Alice Wonderland",
		Age:      20,
		Email:    "alice@example.com",
		GPA:      3.5,
	}

	ok, err := student.Validate()

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}

// 2) Fullname เป็นค่าว่าง -> ต้องได้ error "Fullname is required"
func TestStudent_FullnameRequired(t *testing.T) {
	g := NewGomegaWithT(t)

	student := Student{
		Fullname: "", // ค่าว่าง
		Age:      20,
		Email:    "test@example.com",
		GPA:      3.0,
	}

	ok, err := student.Validate()

	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	// g.Expect(err).To(BeNil())
	g.Expect(err.Error()).To(Equal("Fullname is required"))
}

// 3) อายุ < 18 -> ต้องได้ error "Age must be at least 18"
func TestStudent_AgeMustBeAtLeast18(t *testing.T) {
	g := NewGomegaWithT(t)

	student := Student{
		Fullname: "Bob",
		Age:      16, // อายุ น้อยกว่า 18
		Email:    "bob@example.com",
		GPA:      2.5,
	}

	ok, err := student.Validate()

	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Age must be at least 18"))
}

// 4) Email รูปแบบไม่ถูกต้อง -> ต้องได้ error "Email is invalid"
func TestStudent_EmailInvalidFormat(t *testing.T) {
	g := NewGomegaWithT(t)

	student := Student{
		Fullname: "Charlie",
		Age:      22,
		Email:    "invalid-email-format",
		GPA:      3.0,
	}

	ok, err := student.Validate()

	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Email is invalid"))
}

// 5) GPA ไม่อยู่ในช่วง 0.00-4.00 -> ต้องได้ error "GPA must be between 0.00 and 4.00"
func TestStudent_GPAOutOfRange(t *testing.T) {
	g := NewGomegaWithT(t)

	student := Student{
		Fullname: "David",
		Age:      20,
		Email:    "david@example.com",
		GPA:      4.5, // invalid
	}

	ok, err := student.Validate()

	g.Expect(ok).To(BeFalse())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("GPA must be between 0.00 and 4.00"))
}
