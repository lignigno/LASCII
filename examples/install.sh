#!/bin/bash

# Список исключаемых папок
EXCLUDE="example fonts"

# Перебираем все папки в репозитории
for dir in */; do
	# Если папка не в списке исключений, удаляем её
	if [[ ! " $EXCLUDE " =~ " $dir " ]]; then
		echo "fail: $dir"
		rm -rf "$dir"
	fi
done

echo "complete!"