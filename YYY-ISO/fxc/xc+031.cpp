/*
---------------------------------------
# ✅ xc+031
# ✅ C++ criado: 2021/03/05
# ✅ Objetivo:sql server - C++ Lib
---------------------------------------*/

create proc [dbo].[xc+031]
AS
/*
//TAD - TIPO ABSTRATO DE DADOS

#include <stdio.h>
#include <stdlib.h>

int const MAXTAM = 1000;

int Pilha[MAXTAM];
int Topo;

void Pilha_Construtor(){
    Topo = -1;
}

bool Pilha_Vazia(){
    if(Topo==-1)
        return true;
    else
        return false;

    //return Topo==-1;
}

int Pilha_Tamanho(){
    return Topo+1;
}

bool Pilha_Cheia(){
    if(Topo==MAXTAM-1)
        return true;
    else
        return false;
}

bool Pilha_Push(int valor){ //Push = Empilhar
    if( Pilha_Cheia() ){
       return false;
    }else{
        Topo++;
        Pilha[Topo] = valor;
        return true;
    }
}

bool Pilha_Pop(int &valor){ //Pop = Desempilhar

    if( Pilha_Vazia() ){
        return false;
    }else{
        valor=Pilha[Topo];
        Topo--;
        return true;
    }
}

bool Pilha_Get(int &valor){ //Consulta
    if( Pilha_Vazia() ){
        return false;
    }else{
        valor=Pilha[Topo];
        return true;
    }
}

int main(){

    int valor;

    Pilha_Construtor();
    Pilha_Push(15);
    Pilha_Push(17);
    Pilha_Push(19);
    Pilha_Push(11);
    Pilha_Push(12);

    printf("\n\n Quant. itens da pilha:%d \n\n\n", Pilha_Tamanho() );
    while( Pilha_Pop(valor) ){
        printf("\n %d \n",valor);
    }
    printf("\n\n Quant. itens da pilha:%d \n\n\n", Pilha_Tamanho() );

    Pilha_Construtor();
    Pilha_Push(1945);
    Pilha_Push(1983);

  //  Pilha_Get(valor);
  //  printf("\n %d \n",valor);
    printf("\n\n Quant. itens da pilha:%d \n\n\n", Pilha_Tamanho() );
    while( Pilha_Pop(valor) ){
       printf("\n %d \n",valor);
    }
    printf("\n\n Quant. itens da pilha:%d \n\n\n", Pilha_Tamanho() );

    system("pause");
    return 0;
}


*/