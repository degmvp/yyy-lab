/*
---------------------------------------
# ✅ xc+201 C++ Lib 11
# ✅ C++ criado: 2021/04/14
# ✅ Objetivo:sql server - C++ Lib 11
---------------------------------------*/

create proc [dbo].[xc+201]
AS
/*
#include <iostream>

using namespace std;

class abstractEmployee {
    virtual void askforpromotion()=0;
};

class employee:abstractEmployee {
private:
     string Name;
     string company;
     int Age;
public:
     void setname(string name)  //setter
{
    Name = name;
}
     string getname() {         //getter
        return Name;
     }

     void setcompany(string company)  //setter
{
    company = company;
}
     string getcompany() {         //getter
        return company;
     }
    void setage(int age)  //setter
{
    if(age<=77)
    Age = age;
}
     int getage() {         //getter
        return Age;
     }

     void imp(){
        cout << "Name...... - " << Name << endl;
        cout << "Company... - " << company << endl;
        cout << "Age....... - " << Age << endl << endl;
     }
     employee(string name, string comp, int age){
        Name = name;
        company = comp;
        Age = age;
     }
     void askforpromotion(){
        if(Age > 30)
            cout << Name << " got promoted!\n" << endl;
        else
            cout << Name  << " sorry No promotion for you!\n" << endl;
     }
};
class Developer: employee {
public:
    string FavPro;
    Developer(string nome, string comp, int age, string favpro)
        :employee(nome,comp,age)
    {
         FavPro = favpro;

    }
    void fixbug() {
             cout  << getcompany() <<   " fixed bug usimg " << FavPro << endl << endl;
             cout  << getage()     <<   " fixed bug usimg " << FavPro << endl << endl;
             cout  << getname()    <<   " fixed bug usimg " << FavPro << endl << endl;
             }
};

int main()
{
    employee employ1 = employee("Saldina", "YT-Code",25);
    employ1.imp();

    employee employ2 = employee("Degmar", "C++ OOP-Code",75);
    employ2.imp();


 //   employ1.setage(39);
    employ2.setage(76);
    cout << employ1.getname() <<" is " <<employ1.getage() << " years old" << endl;
    cout << employ2.getname() <<" is " <<employ2.getage() << " years old" << endl << endl;

    employ1.askforpromotion();
    employ2.askforpromotion();

    Developer d = Developer("Degmar", "Programador",75,"C++");
    d.fixbug();


    system("pause");
    return 0;
}


*/