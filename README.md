# 📡 TCP CHAT - Golang

## 📌 Description
TCP CHAT est une application de communication en terminal développée en **Go** utilisant le protocole **TCP/IP**.

Le projet permet à deux utilisateurs de communiquer en temps réel via :
- un serveur
- un client

L’objectif principal de ce projet était de découvrir :
- les sockets TCP
- la communication réseau
- les goroutines
- la concurrence en Go

---

## 🛠️ Technologies utilisées
- Go (Golang)
- TCP/IP
- Terminal / CLI

---

## 🚀 Fonctionnalités
- Communication en temps réel
- Architecture client / serveur
- Envoi et réception simultanés des messages
- Utilisation des goroutines pour le multitâche
- Interface terminal simple et légère

---

## 📁 Structure du projet

```text
TCP-CHAT/
├── main.go
├── server.go
└── client.go
```

---

## ⚙️ Installation

### 1. Cloner le projet

```bash
git clone https://github.com/ton-username/TCP-CHAT.git
cd TCP-CHAT
```

### 2. Initialiser les modules Go

```bash
go mod init tcp-chat
go mod tidy
```

---

## ▶️ Utilisation

### Lancer le serveur

```bash
go run . server
```

Le serveur écoute sur le port :

```text
:8080
```

---

### Lancer le client

```bash
go run . client
```

Le client se connecte à l’adresse IP définie dans `client.go` :

```go
net.Dial("tcp", "10.36.0.35:8080")
```

⚠️ Pensez à modifier l’adresse IP selon votre réseau local.

---

## 🧠 Fonctionnement technique

### Serveur
Le serveur :
- ouvre un port TCP
- attend une connexion
- reçoit les messages
- envoie les messages

Fonctions utilisées :
- `net.Listen()`
- `Accept()`
- `goroutine`

---

### Client
Le client :
- se connecte au serveur
- lit les messages reçus
- envoie les messages écrits dans le terminal

Fonctions utilisées :
- `net.Dial()`
- `bufio.NewReader()`
- `bufio.NewScanner()`

---

## 📚 Concepts appris
- Réseau TCP/IP
- Communication client / serveur
- Concurrence avec les goroutines
- Gestion des entrées/sorties
- Lecture et écriture sur une connexion TCP

---

## 🔧 Améliorations possibles
- Gestion de plusieurs clients
- Ajout de pseudonymes
- Historique des messages
- Interface graphique
- Chiffrement des communications
- Commandes personnalisées (`/quit`, `/help`, etc.)

---

## 👤 Auteur
Enzo Courvalet

---

## 📄 Licence
Projet réalisé dans un cadre pédagogique.