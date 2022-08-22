
[![CircleCI](https://dl.circleci.com/status-badge/img/gh/abaron10/Gothon/tree/master.svg?style=shield)](https://dl.circleci.com/status-badge/redirect/gh/abaron10/Gothon/tree/master)
<!-- PROJECT LOGO -->
<br />
<p align="center">
 <img src="https://github.com/abaron10/Gothon/blob/master/static/Gothon.png?raw=true)" />
  <h3 align="center">Gothon library</h3>
  
</p>

The main goal of the Gothon's library is to facilitate and speed up the development process by taking inspiration from some cool methods Python coding language. With the version of Go 1.18 the use of generics makes it much easier to create generic methods with a smaller amount of code, this library takes advantage of the best of this new feature

<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About the project</a>
      <ul>
        <li><a href="#usage">Usage</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
       <ul>
        <li><a href="#supported-methods">Supported Methods</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgements">Acknowledgements</a></li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About The Project

### Usage

The following example shows the use of the extends method, which adds two lists and transform them into a single list.
As said before, this library works with Go 1.18, so with the use of generics you don't need to call another function to achieve the same goal with slices of other types. Reference to Python usage https://www.w3schools.com/python/ref_list_extend.asp

<p align="center">
 <img src="https://github.com/abaron10/Gothon/blob/master/static/demo2.png?raw=true)" />
</p>


This second example makes use of the Pop() function supporting negative indexing. In python a -1 in a list corresponds to position len(n) -1, a -2 to position len(n) -2 and so on.

<p align="center">
 <img src="https://github.com/abaron10/Gothon/blob/master/static/demo1.png?raw=true)" />
</p>

<!-- GETTING STARTED -->
## Getting Started
To install the library in your project, do the following:
```
go get github.com/abaron10/Gothon

```
Please remember this library only works with Go ˆ1.18


### Supported Methods

The following methods and data structures are supported for the library in version v0.0.1

## Slices

| Merthod Name      | Description |
| ----------- | ----------- |
|  Extend     | This method adds the specified list elements to the end of the current list. Requires two slices and returns a new slice with the extesion of both|
| Pop   | This method removes the element at the specified position, It requires a pointer to the slice and the index to be removed. Changes are made in place with the given slice and the method returns the popped value. Supports negative indexing e.i -1 corresponds to the last position of a slice  |
|Index |This method returns the position at the first occurrence of the specified value.It requires an slice and a value to be checked inside the slice. Returns an integer with the index of the element, if the value doesn't exist return -1|
|Contains|This method check if the item exists in the slice. It requires an slice and a value to be checked inside the slice. Returns a boolean whether the element exists or not|
|Reverse|This method reverses the sorting order of the elements. Receives a pointer from the given slice. Changes are made in place with the given slice.|
|Remove|This method removes the first occurrence of the element with the specified value. It requires a pointer to the slice and the raw value to be removed from the slice. Changes are made in place with the given slice|
|Insert|This method inserts the specified value at the specified position. It requires a pointer to the slice, an index, and a value to be inserted. Changes are made in place with the given slice|
|Clear|This method removes all the elements from the slice|
|Count|This method returns the number of elements with the specified passed value|
|Copy|This method returns a copy of the specified slice|

### Installation


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch 
3. Commit your Changes 
4. Push to the Branch 
5. Open a Pull Request

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->
## Contact

LinkedIn: [Andrés Barón Sandoval](https://www.linkedin.com/in/andres-baron-sandoval-76ab96186/)


Personal Contact Page: [https://abaron10.github.io/Andres_Baron_contact_page/index.html](https://abaron10.github.io/Andres_Baron_contact_page/index.html)

Project Link: [https://github.com/abaron10/Gothon](https://github.com/abaron10/Gothon)

<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements
* [Go Generics](https://go.dev/doc/tutorial/generics)
* [Python Data structures](https://developers.google.com/edu/python/lists)
* [Choose an Open Source License](https://choosealicense.com)
* [GitHub Pages](https://pages.github.com)








