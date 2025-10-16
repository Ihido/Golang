# 🔥 Git команды для контроля версий

## 🚀 Быстрая настройка нового проекта с GitHub
```bash
# Инициализация и первый коммит
git init
git add .
git commit -m "Первоначальный коммит"

# Связь с GitHub репозиторием
git remote add origin https://github.com/username/repository.git
git branch -M main
git push -u origin main
```

## 📦 Основные команды (использую каждый день)
```bash
git status                      # Показать изменения
git add .                       # Добавить ВСЕ файлы
git add имя_файла               # Добавить конкретный файл
git commit -m "Сообщение"       # Сделать коммит
git push                        # Отправить на GitHub
git pull                        # Скачать изменения
```

## 📝 Работа с изменениями
```bash
git diff                        # Показать изменения в файлах
git diff --staged               # Показать изменения в индексе
git restore файл                # Отменить изменения в файле
git restore --staged файл       # Убрать файл из индекса
git commit --amend              # Исправить последний коммит
```

## 🌿 Ветки (branch)
```bash
git branch                      # Показать все ветки
git branch имя_ветки            # Создать ветку
git checkout имя_ветки          # Перейти в ветку
git checkout -b новая_ветка     # Создать и перейти в ветку
git merge ветка                 # Объединить ветки
git branch -d ветка             # Удалить ветку
```

## 📚 История и логи
```bash
git log                         # Показать историю коммитов
git log --oneline               # Краткая история
git log --graph                 # История с графиком веток
git show хеш_коммита            # Показать конкретный коммит
git blame файл                  # Кто и когда менял файл
```

## 🚀 Работа с GitHub
```bash
git clone ссылка_репозитория    # Скачать репозиторий
git remote -v                   # Показать связанные репозитории
git push origin main            # Отправить в ветку main
git push -u origin ветка        # Отправить новую ветку
git fetch                       # Скачать информацию об изменениях
```

## 🛠 Полезные сценарии

### Обновление репозитория после изменений
```bash
git add .
git commit -m "Описание изменений"
git push
```

### Проверка статуса перед коммитом
```bash
git status
git diff
```

## ⚠️ Решение частых ошибок

### Если "remote origin already exists"
```bash
git remote remove origin
git remote add origin https://github.com/username/repository.git
```

### Если "failed to push some refs"
```bash
git pull origin main --allow-unrelated-histories
git push -u origin main
```

### Если "authentication failed"
1. Зайди в GitHub → Settings → Developer settings → Personal access tokens
2. Создай новый токен с правами `repo`
3. Используй токен вместо пароля

## ❌ Что делать если...
```bash
# Случайно сделал коммит не в ту ветку
git checkout правильная_ветка
git cherry-pick хеш_коммита

# Забыл добавить файл в последний коммит
git add забытый_файл
git commit --amend --no-edit

# Хочу отменить последний коммит (локально)
git reset --soft HEAD~1
```

## 💾 Временное сохранение изменений
```bash
git stash                       # Временно сохранить изменения
git stash pop                   # Вернуть сохраненные изменения
git stash list                  # Показать все сохранения
```

## 🏷️ Работа с тегами версий
```bash
git tag v1.0.0                  # Создать тег версии
git push --tags                 # Отправить теги на GitHub
```

---

**💡 Совет:** Используй `git status` перед каждым коммитом чтобы убедиться что все файлы добавлены!