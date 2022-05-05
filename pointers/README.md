# Pointer

**Concept**  
Variables are the names given to a memory location where the actual data is stored. To access the stored data we need the address of that particular memory location. To remember all the memory addresses(Hexadecimal Format) manually is an overhead that’s why we use variables to store data and variables can be accessed just by using their name.  

A pointer is a special kind of variable that is not only used to store the memory addresses of other variables but also points where the memory is located and provides ways to find out the value stored at that memory location. It is generally termed as a Special kind of Variable because it is almost declared as a variable but with *(dereferencing operator).  

Reference.  
https://www.geeksforgeeks.org/pointers-in-golang/  
https://www.youtube.com/watch?v=sTFJtxJXkaY

## สรุป  

**Variable** คือตัวเก็บ address ของ value ที่อยู่ใน memory  
**Pointer** นั้นคล้ายกับ variable โดยตัวมันก็เก็บ address ของข้อมูลของมันที่อยู่ใน memory แต่มันพิเศษตรงที่ตัวข้อมูลที่มันเก็บอยู่นั้นไม่ใช่ value ปกติ แต่เป็น address ของ value อื่นที่อยู่ใน memory

Variable (i, j)
![image](https://user-images.githubusercontent.com/50371660/166864766-9ae5763c-db4f-4ce4-ae3f-9fc29f19ada5.png)  

Variable (i, j) and Pointer (p)
![image](https://user-images.githubusercontent.com/50371660/166866775-f1b3adf9-f42b-4d5c-aac6-24b9e378d134.png)

![image](https://user-images.githubusercontent.com/50371660/166713294-c2cf031a-0666-4b04-aa4e-0c3321151127.png)
