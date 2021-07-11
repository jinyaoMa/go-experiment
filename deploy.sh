# abort on errors
#set -e

# navigate into the build output directory
cd "view/dist"

# if you are deploying to a custom domain
echo 'go-experiment.ma-jinyao.cn' > CNAME

git init
git add -A
git commit -m 'deploy'

# if you are deploying to https://<USERNAME>.github.io/<REPO>
git push -f git@github.com:jinyaoMa/go-experiment.git main:gh-pages

cd -