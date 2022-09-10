/*
MIT License

Copyright (c) 2022 Jugesuke

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

/*
Package dajareGo provides checking a sentence if it is Dajare (Japanese pun).

# Definition of Dajare

In this Package, Dajare is defined the sentence contains pair(s) of words which has a similar reading but a different meaning.

# How to Use

You can check if a sentence is Dajare with IsDajare function, like below.

	if err := dajareGo.Init(); err != nil {
		panic(err)
	}
	result := dajareGo.IsDajare("アルミ缶の上にあるミカン")
	if result.IsDajare {
		fmt.Println("This is Dajare")
	} else {
		fmt.Println("This is not Dajare")
	}
*/
package dajareGo
