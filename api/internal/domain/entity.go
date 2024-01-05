package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"regexp"
	"strconv"
	"time"
)

const (
	USER_ID_PATTERN = `^[a-zA-Z0-9]{10}$`
	USER_EMAIL_PATTERN = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	POSTAL_CODE_PATTERN = `^\d{3}-\d{4}$`
	PHONE_NUMBER_PATTERN = `^\d{3}-\d{4}-\d{4}$`
)

var (
	INVALID_USER_ID = errors.New("invalid user id")
	INVALID_USER_EMAIL_ADDRESS = errors.New("invalid user email address")
	INVALID_FIRST_NAME = errors.New("invalid user first name")
	INVALID_LAST_NAME = errors.New("invalid user last name")
	INVALID_PHONE_NUMBER = errors.New("invalid phone number")
	INVALID_POST_OFFICE_NUMBER = errors.New("invalid post office number")
	INVALID_BIRTH_DAY = errors.New("invalid birth day")
	INVALID_GENDER = errors.New("invalid gender")
	INVALID_PREF = errors.New("invalid pref")
	INVALID_CITY = errors.New("invalid city")
	INVALID_EXTRA = errors.New("invalid extra")
	INVALID_PASSWORD = errors.New("invalid password")
)

// ドメインモデル
type user struct {
	id               userUUID         
	email            userEmail        
	password         userPassWord     
	user_id          userID           
	firstName        userFirstName    
	lastName         userLastName     
	gender           userGender       
	birthDay         userBirthDay     
	phoneNumber      userPhoneNumber  
	postOfficeNumber postOfficeNumber 
	pref             pref             
	city             city             
	extra            extra            
}

// ドメイン バリューオブジェクト
type userUUID struct{ value string }
type userID struct{ value string }
type userEmail struct{ value string }
type userPassWord struct{ value string }
type userFirstName struct{ value string }
type userLastName struct{ value string }
type userPhoneNumber struct{ value string }
type userGender struct{ value string }
type userBirthDay struct{ value string }
type postOfficeNumber struct{ value string }
type pref struct{ value string }
type city struct{ value string }
type extra struct{ value string }

// ドメインルール

/*
userID バリデーション godoc
* 10文字
* 英数字
* 記号なし
*/
func(v *userID) valid() error {
    r := regexp.MustCompile(USER_ID_PATTERN)
    matched := r.MatchString(v.value)

    // 結果を出力
    if !matched {
        return INVALID_USER_ID
    }

	return nil
}

/* userEmail バリデーション godoc メールアドレスの形式のなっていること */
func(v *userEmail) Valid() error {
	match, _ := regexp.MatchString(USER_EMAIL_PATTERN, v.value)
	if !match {
		return INVALID_USER_EMAIL_ADDRESS
	}

	return nil
}

/* userEmail バリデーション godoc 1文字以上*/
func(v *userFirstName) Valid() error {
	if v.value == "" {
		return INVALID_FIRST_NAME
	}

	return nil
}

/* userLastName バリデーション godoc 1文字以上 */
func(v *userLastName) Valid() error {
	if v.value == "" {
		return INVALID_LAST_NAME
	}

	return nil
}

/* userBirthDay バリデーション godoc RFC3339形式であること */
func(v *userBirthDay) Valid() error {
	t := v.value + "T00:00:00Z"
	_, err := time.Parse(time.RFC3339, t)
	return err
}

/* userGender バリデーション godoc 1:男、2:女、3:その他であること */
func(v *userGender) Valid() error {
	g, _ := strconv.Atoi(v.value)
	if g < 1 && g > 3 {
		return INVALID_GENDER
	}

	return nil
}

/* postOfficeNumber バリデーション godoc 郵便番号の形式であること */
func(v *postOfficeNumber) Valid() error {
	match, _ := regexp.MatchString(POSTAL_CODE_PATTERN, v.value)
	if !match {
		return INVALID_POST_OFFICE_NUMBER
	}

	return nil
}

/* userPhoneNumber バリデーション godoc 電話番号の形式であること */
func(v *userPhoneNumber) Valid() error {
	match, _ := regexp.MatchString(PHONE_NUMBER_PATTERN,v.value)
	if !match {
		return INVALID_PHONE_NUMBER
	}

	return nil
}

/* pref バリデーション godoc 1文字以上であること */
func(v *pref) Valid() error {
	if v.value == "" {
		return INVALID_PREF
	}

	return nil
}

/* city バリデーション godoc 1文字以上であること */
func(v *city) Valid() error {
	if v.value == "" {
		return INVALID_CITY
	}

	return nil
}

/* extra バリデーション godoc 1文字以上であること */
func(v *extra) Valid() error {
	if v.value == "" {
		return INVALID_EXTRA
	}

	return nil
}

/* userPassWord godoc パスワードのハッシュ化 */
func(v *userPassWord) tohash(sault string) string {
	hasher := sha256.New()
	hasher.Write([]byte(v.value + sault))
	hashByte := hasher.Sum(nil)

	hash := hex.EncodeToString(hashByte)
	return hash
}

/* userPassWord godoc パスワードのデコード */
func(v *userPassWord) decode(hash string, sault string) error {
	hasher := sha256.New()
	hasher.Write([]byte(v.value + sault))
	hashByte := hasher.Sum(nil)

	encodeHash := hex.EncodeToString(hashByte)

	if encodeHash == hash {
		return INVALID_PASSWORD
	}
	return nil
}

// バリューオブジェクトの取得関数
func (u *userUUID) GetUUID() string {return u.value}
func (u *userID) GetID() string {return u.value}
func (u *userEmail) GetEmail() string {return u.value}
func (u *userFirstName) GetFirstName() string {return u.value}
func (u *userLastName) GetLastName() string {return u.value}
func (u *userGender) GetGender() string {return u.value}
func (u *userBirthDay) GetBirthDay() string {return u.value}
func (u *userPassWord) GetPassWord() string {return u.value}
func (u *userPhoneNumber) GetPhoneNumber() string {return u.value}
func (u *postOfficeNumber) GetPostOfficeNumber() string {return u.value}
func (u *pref) GetPref() string {return u.value}
func (u *city) GetCity() string {return u.value}
func (u *extra) GetExtra() string {return u.value}
