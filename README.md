# Pinpin

🇫🇷

Pinpin est un client **non officiel** pour l’enceinte **Merlin**, développé en
Go. Il permet de téléverser directement sur l’enceinte vos **fichiers audio
personnels** (musique, histoires, podcasts libres de droits, etc.).

Nécessite :
- Go 1.24.3+
- ffmpeg

## Utilisation

```bash
go install github.com/gawen/pinpin/cmd/pinpin@latest
```

```bash
# Créer un dossier qui contiendra votre arborescence de fichiers Pinpin.
$ mkdir mes_fichiers_pinpin
$ cd mes_fichiers_pinpin

# Créer les dossiers qui apparaîtront dans le Merlin.
$ mkdir Historias
$ cd Historias

# Placez-y vos histoires, dans un format (audio ou vidéo) que ffmpeg peut lire.
# Le nom des fichiers sera le nom des histoires qui apparaîtront dans le
# Merlin.
$ ls
Gato.mp3
Hormiga.webm
Mariposa.mp4
Mariquita.webm
Perro.mp3

# Optionnellement, placez-y les images qui seront utilisées en tant que miniatures
# pour vos histoires. Pour qu'une image soit utilisée, celle-ci doit avoir exactement
# le même nom que les fichiers audio/vidéo de vos histoires avec l'extension .jpg
# L'image utilisée doit être carré et faire 128px de côté
$ ls
Gato.mp3
Gato.jpg

# Enfin, exécutez Pinpin et suivez les instructions.
$ cd ../../
$ ${GOPATH}/bin/pinpin mes_fichiers_pinpin
```

## Légal

Veuillez lire le fichier [`DISCLAIMER.md`](DISCLAIMER.md).
