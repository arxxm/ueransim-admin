Получение списка доступных объектов:
- GET host:8080/api/v1/base-stations/ - Получение списка доступных в эмуляторе базовых станций
- GET host:8080/api/v1/mobile-terminals/ - Получение списка доступных в эмуляторе UE абонентов

Примеры получения данных:
- nr-gnb:
  - GET host:8080/api/v1/base-stations/cc     - Количество подключенных UE
  - GET host:8080/api/v1/base-stations/info   - Получение списка доступных в эмуляторе базовых станций
  - GET host:8080/api/v1/base-stations/status - Статус БС

- nr-ue:
  - GET host:8080/api/v1/mobile-terminals/status - Статус мобильного терминала
  - GET host:8080/api/v1/mobile-terminals/cs - Статус подключения к сети передачи данных