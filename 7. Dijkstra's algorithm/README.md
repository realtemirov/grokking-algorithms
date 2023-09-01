---
cover: >-
  https://images.unsplash.com/photo-1586527155314-1d25428324ff?crop=entropy&cs=srgb&fm=jpg&ixid=M3wxOTcwMjR8MHwxfHNlYXJjaHwxMHx8dGhpbmt8ZW58MHx8fHwxNjkzNDQyMzU5fDA&ixlib=rb-4.0.3&q=85
coverY: 0
layout:
  cover:
    visible: true
    size: full
  title:
    visible: true
  description:
    visible: true
  tableOfContents:
    visible: true
  outline:
    visible: true
  pagination:
    visible: true
---

# Dijkstra's algorithm

• Biz grafiklarni muhokama qilishni davom ettiramiz va siz vaznli grafiklar haqida bilib olasiz: ba'zi qirralarga ko'proq yoki kamroq og'irlik berish usuli.

• Siz Dijkstra algoritmini o'rganasiz, bu sizga "Xga eng qisqa yo'l nima?"ss degan savolga javob berishga imkon beradi. vaznli grafiklar uchun.

• Siz Dijkstra algoritmi ishlamaydigan grafiklardagi sikllar haqida bilib olasiz.

Oxirgi bobda siz A nuqtadan B nuqtaga o'tish yo'lini aniqladingiz.

![point](image.png)

Bu eng tezkor yo'l bo'lishi shart emas. Bu eng qisqa yo'l, chunki u eng kam sonli segmentlarga ega (uch segment). Ammo siz ushbu segmentlarga sayohat vaqtlarini qo'shsangiz, deylik. Endi siz tezroq yo'l borligini ko'rasiz.

![times](image-1.png)

Oxirgi bobda birinchi boʻlib qidiruvdan foydalangansiz. Kenglik-birinchi qidiruv sizga eng kam segmentli yo'lni topadi (birinchi grafik bu erda ko'rsatilgan). Buning o'rniga eng tez yo'lni xohlasangiz nima bo'ladi (ikkinchi grafik)? Buni `Dijkstra algoritmi` deb ataladigan boshqa algoritm yordamida `eng tez` bajarishingiz mumkin.