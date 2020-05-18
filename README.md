# Как сейчас всё устроено

Firebase.

Используются Authentication, Cloud Firestore, Functions, Storage

## Movies

Имеется одна коллекция -- `movies`. [Пример документа из этой коллекции](#movie)

### add movie

что учесть: фильм создается вместе с 9 скриншотами

как выполнено:

1. [Найти фильм из themoviedb.org](#search-themoviedborg) (возвращается объект `movie`, без [screens](#screens))
1. Загрузка скринов (возвращается массив [screens](#screens) с названиями файлов и его порядком)
   - Загрузить скрины на firestore storage
   - Firestore Functions отлавливает загрузку скринов и вызывает ImageResizer. Она уменьшает пикчу и сохраняет с таким названием: `<имя_файла>_500x500`
1. Создать документ `movie`:
   - Взять объект из первого шага и задать ему ключ `screens` со значениями из второго шага.
   - Записать полученный объект в коллекцию `movies`
   - Firestore Functions отлавливает создание документа в коллекции и увеличивает счетчик в документе `--stats--`

### delete movie

Достаточно удалить документ в `movies` и затем Firestore Functions отловит это событие и удалит принадлежащие ему файлы и уменьшит счетчик в `--stats--`

## Users

С юзерами проделана самая минимальная работа. На данном этапе этого хватает для отладки.

Создаю пользователя в консоли.

## Референсы

### Movie

[Назад к "Movies"](#movies)

```js
{
  id: 'GEuMOfz7SImnBMV76CE4',
  adult: false,
  backdrop_path: '/5gPPx16QWx071VAI1M0RAVKJ6tc.jpg',
  belongs_to_collection: {
    backdrop_path: '/zJd4E0ImT1U6O8k6aX3ssc9iUBS.jpg',
    id: 295,
    name: 'Pirates of the Caribbean Collection',
    poster_path: '/zRBaZxS5YauLvRYjAdL4AUCwlht.jpg'
  },
  budget: 200000000,
  createdAt: -1586499559082,
  directLink: 'Pirates_of_the_Caribbean__Dead_Man_s_Chest_58',
  genres: [
    { id: 12, name: 'Adventure' },
    { id: 14, name: 'Fantasy' },
    { id: 28, name: 'Action' }
  ],
  homepage: 'http://disney.go.com/disneypictures/pirates/',
  imdb_id: 'tt0383574',
  name: "Pirates of the Caribbean: Dead Man's Chest",
  original_language: 'en',
  original_title: "Pirates of the Caribbean: Dead Man's Chest",
  overview: 'Captain Jack Sparrow works his way out of a blood debt with the ghostly Davey Jones, he also attempts to avoid eternal damnation.',
  popularity: 52.451,
  poster_path: '/AdRQGfT05z6L9gIpUpkh4McMmpm.jpg',
  production_companies: [
    {
      id: 2,
      logo_path: '/wdrCwmRnLFJhEoH8GSfymY85KHT.png',
      name: 'Walt Disney Pictures',
      origin_country: 'US'
    },
    {
      id: 130,
      logo_path: '/c9dVHPOL3cqCr2593Ahk0nEKTEM.png',
      name: 'Jerry Bruckheimer Films',
      origin_country: 'US'
    },
    {
      id: 19936,
      logo_path: null,
      name: 'Second Mate Productions',
      origin_country: 'US'
    }
  ],
  production_countries: [ { iso_3166_1: 'US', name: 'United States of America' } ],
  release_date: '2006-06-20',
  revenue: 1065659812,
  runtime: 151,
  screens: [
    {
      name: string, // имя должно быть уникальным (как делаю сейчас: `id` фильма из themovie.org + рандом uuid)
      order: number, // порядок даю из фронта
      publicUrls: {
        thumb: `https://firebasestorage.googleapis.com/v0/b/ahoy-kino.appspot.com/o/screens%2Fthumbs%2F${fileName}_500x500?alt=media`,
        full: `https://firebasestorage.googleapis.com/v0/b/ahoy-kino.appspot.com/o/screens%2F${fileName}?alt=media`,
      },
    },
    ...
  ],
  spoken_languages: [
    { iso_639_1: 'en', name: 'English' },
    { iso_639_1: 'tr', name: 'Türkçe' },
    { iso_639_1: 'el', name: 'ελληνικά' },
    { iso_639_1: 'zh', name: '普通话' }
  ],
  status: 'Released',
  tagline: 'Jack is back!',
  title: "Pirates of the Caribbean: Dead Man's Chest",
  tmdb_id: 58,
  userUid: 'DnEafM9l3IPdMJarcOi3YGNfKIm2',
  video: false,
  vote_average: 7.2,
  vote_count: 10641
}
```

### Screens

[Назад к "add movie"](#add-movie)

```js
[
  {
    name: string, // имя должно быть уникальным (как делаю сейчас: `id` фильма из themovie.org + рандом uuid)
    order: number, // порядок даю из фронта
    publicUrls: {
      thumb: `https://firebasestorage.googleapis.com/v0/b/ahoy-kino.appspot.com/o/screens%2Fthumbs%2F${fileName}_500x500?alt=media`,
      full: `https://firebasestorage.googleapis.com/v0/b/ahoy-kino.appspot.com/o/screens%2F${fileName}?alt=media`,
    },
  },
  ...
]
```

### Search themoviedb.org

[Назад к add movie](#add-movie)

АПИ для поиска: `https://api.themoviedb.org/3/search/multi?api_key=${apiKey}&language=en-US&query=${query}&page=1&include_adult=false`

[документация themoviedb.org](https://developers.themoviedb.org/3/getting-started/introduction)

ключ: f0844452c473688f4cc95d1d40e4a1d8
