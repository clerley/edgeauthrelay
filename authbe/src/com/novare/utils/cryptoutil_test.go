package utils

import "testing"

/*
MIT License

Copyright (c) 2020 Clerley Silveira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

func TestCryptoUtils(t *testing.T) {

	pass := "123456"
	if IsSecurePassword(pass) {
		t.Errorf("The password is week but, it was marked as secure")
		return
	}

	hPass, ok := GetPassword("123456", "SALT")
	if !ok {
		t.Errorf("The password was not generated")
		return
	}

	if !IsValidPassword("123456", "SALT", hPass) {
		t.Error("The function IsValidPassword replied with false!")
		return
	}

	if IsValidPassword("654321", "SALT", hPass) {
		t.Error("The function IsValidPassword replied with true!")
		return
	}

	if IsValidPassword("123456", "PEPPER", hPass) {
		t.Error("The function IsValidPassword replied with true! Condition:2 - SALT is not the same")
		return
	}

}
