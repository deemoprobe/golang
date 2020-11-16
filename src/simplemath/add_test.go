/*
 * @Author: deemoprobe@gmail.com
 * @Date:   2020-11-16 16:13:36
 * @Last Modified by:   deemoprobe@gmail.com
 * @Last Modified time: 2020-11-16 16:17:41
 */
// Packeage simplemath clac
package simplemath

import "testing"

func TestAdd1(t *testing.T) {
	r := Add(1, 2)
	if r != 3 {
		t.Errorf("Add(1, 2) failed. Got %v, expected 3.", t)
	}
}
