
---

# JSON Terminology in Go

When working with JSON in Go, two key concepts are essential to understand:

1. **Marshalling**: This is the process of converting a Go data structure into valid JSON format.
2. **Unmarshalling**: This involves parsing a valid JSON string and converting it into a corresponding Go data structure.

In other programming languages, the term "marshalling" is often referred to as "serializing," while "unmarshalling" is commonly known as "deserializing." For clarity, we will use Go's terminology throughout this document.

---

# Pitfalls with JSON Unmarshalling in Go

When unmarshalling JSON in Go, be aware of the following common pitfalls:

1. **Extra Fields Omitted**: Any fields in the JSON that do not correspond to the struct will be ignored.
2. **Zero Value Fallback**: If the JSON is missing fields, the corresponding fields in the struct will default to their zero values.
3. **Case Insensitivity**: Unmarshalling is case insensitive, which may lead to unexpected behavior.
4. **Exact Field Name Matching**: Field names in the struct must match the JSON keys exactly, including case.

---

# Customizing JSON Field Names with Struct Tags

In Go, struct tags are annotations that provide additional information about how the fields should be treated by various libraries and tools. These tags are strings that are added to the end of a field declaration, enclosed in backticks.

A common use case for struct tags is to customize how a struct is marshalled and unmarshalled to and from JSON. By adding tags to struct fields, you can control field naming and validation.

---

# How to Run the Code

To test the code, use the following command:

```bash
go run golang/go-json/main.go
```

--- 
