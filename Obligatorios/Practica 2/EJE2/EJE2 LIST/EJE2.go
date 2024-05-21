package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
	"container/list"
)

type Transaccion struct{
	Monto float64
	IDEnvio string
	IDRecibo string
	Fecha string
}

type Block struct{
	Hash string
	HashPrev string
	Data Transaccion
	Fecha string
}

type Wallet struct{
	ID string
	Nombre string
	Apellido string
}

type Blockchain struct{
	Blocks *list.List
}

// LIST

func createWallet(nombre,apellido string) Wallet {
	datos:= nombre + apellido + time.Now().String()
	id:= fmt.Sprintf("%x",sha256.Sum256([]byte(datos)))
	return Wallet{
		ID: id,
		Nombre: nombre,
		Apellido: apellido,
	}
}

func sendTransaction(blockchain *Blockchain, senderId,reciverId string, monto float64)  {
	fecha := time.Now().Format("2006-01-02 15:04:05")
	trans := Transaccion{
		Monto: monto,
		IDEnvio: senderId,
		IDRecibo: reciverId,
		Fecha: fecha,
	}

	var prevHash string
	if (blockchain.Blocks.Len())== 0{
		prevHash = "empty" // si es el primero y el hash anterior no existe pongo en blanco
	}else{
		prevHash = blockchain.Blocks.Back().Value.(Block).Hash // agarro el hash del anterior
	}

	block := Block{
		HashPrev: prevHash,
		Data: trans,
		Fecha: fecha,
	}

	// crear el hash para este block
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%v",block)))
	block.Hash = hex.EncodeToString(hash.Sum(nil))

	// lo meto al block
	blockchain.Blocks.PushBack(block)
}

func insertBroken(bc *Blockchain, pos int)  {
	valid := true
	if bc.Blocks.Len() < pos{
		valid = false
	}
	
}


func main()  {
	bc := &Blockchain{
		Blocks: list.New(),
	}

	wallet1:=createWallet("camilo","herrera")
	wallet2:=createWallet("micaela","d'agostino")
	wallet3:=createWallet("santiago","herrera")
	wallet4:=createWallet("morena","herrera")
	wallet5:=createWallet("tomas","rivas")
	
	sendTransaction(bc,wallet1.ID, wallet2.ID,777.99)
	sendTransaction(bc,wallet4.ID, wallet3.ID,9983.99)
	sendTransaction(bc,wallet5.ID, wallet1.ID,63441.99)
	
	for e := bc.Blocks.Front(); e != nil; e = e.Next() {
		block := e.Value.(Block)
		fmt.Printf("Hash: %s, HashPrev: %s, Monto: %f, IDEnvio: %s, IDRecibo: %s, Fecha: %s\n",
			block.Hash, block.HashPrev, block.Data.Monto, block.Data.IDEnvio, block.Data.IDRecibo, block.Fecha)
		fmt.Println(" ")
	
	}
}