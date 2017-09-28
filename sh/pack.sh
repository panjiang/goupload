#!/bin/sh

set -e

name="upload"
export_dir="/share/"
if [ $# != 0 ]; then
    export_dir=$1
fi

project_home=$PWD
echo "current branch:"
git rev-parse --abbrev-ref HEAD
# make sure
echo -n "version: "
read version

if [ -z $version ]; then
    version=$(date +"%g%m%d%H")
    echo "auto version: $version"
fi


package="${name}_v$version"
dir_name="${name}"
tagname="${name}-$version"
temp_dir=/tmp/$dir_name

if [ $temp_dir = $project_home ]; then
    echo "temp directory cann't be same with project!"
    exit 0
fi

echo -n "package: $package.tar.gz (Y/n):"
read confirm
if [ "$confirm" != "Y" ];then
    echo "abandoned pack"
    exit 0
fi

# copy to temp directory
if [ -d $temp_dir ];then
    rm -rf $temp_dir
fi
echo "mkdir $temp_dir"
mkdir $temp_dir

go build ${name}.go
for dir in $name templates README.md config.json; do
    echo "cp -rf $dir $temp_dir/$dir"
    cp -rf $dir $temp_dir/$dir
done

echo "cd $temp_dir"
cd $temp_dir

if [ $PWD != $temp_dir ]; then
    echo "cd $temp_dir failed!"
    exit 0
fi


# mark version
echo $version > VERSION

cd ../
echo "tar -czf $package.tar.gz $dir_name"
tar -czf $package.tar.gz $dir_name
rm -rf $package
mv $package.tar.gz $project_home/

cd $project_home/
echo "md5 $package.tar.gz"
#md5sum $package.tar.gz > $package.md5

echo "pack $package success"

echo -n "tag: $tagname (Y/n):"
read confirm
if [ "$confirm" != "Y" ];then
    echo "abandoned tag"
    exit 0
fi

git tag -d $tagname  2> /dev/null || true
git tag $tagname
git push origin --delete refs/tags/$tagname
git push origin $tagname:refs/tags/$tagname

echo "tag $tagname success"
exit 0
