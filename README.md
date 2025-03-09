# bolos
bolos
Este é meu arquivo de bolo

git init
git add .
#Configurado o git, fazer apenas 1x
git config --global user.email "alexsjmaia@gmail.com"
git config --global user.name "alexsjmaia"
git commit -m "projeto inicial"
git branch -M main
git remote add origin git@github.com:alexsjmaia/bolos.git
git push -u origin main
ssh-keygen -t ed25519 -C "alexsjmaia@gmail.com"

git init
git add .
git commit -m "projeto inicial"
git branch -M main
git remote add origin git@github.com:alexsjmaia/bolos.git
git push -u origin main

#Se der erro faça o comando abaixo
git pull origin main --rebase
git add .
git rebase --continue
git push origin main

#Ou fazer o merge
git pull origin main
git push origin main


