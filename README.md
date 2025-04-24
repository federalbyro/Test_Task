
## Какие запросы и как принимает данный сервис(вроде бы по ТЗ)

* Refresh Token Endpoint:
```bash
curl -X PUT http://localhost:2518/tokens/refresh \
     -H "Content-Type: application/json" \
     -d '{"refresh_token":"884252f3-adc4-4464-b4f0-aed5a69b3262.bOqqmjB/cA1IP2zCA/PVrYAcdgXADYQ0NudUCoLErWU="}'
```
* Access Token Endpoint:
```bash
$ curl -X POST "http://localhost:2518/tokens/access?GUID=123e4567-e89b-12d3-a456-426614124000"
```



### Первичный макет
<div>
<p>Figma:</p>
</div>

[![Figma](https://img.shields.io/badge/Figma-F24E1E?style=for-the-badge&logo=figma&logoColor=white)](https://www.figma.com/board/dhZQEAmwZDhEFjbGDoSqFh/Untitled?node-id=0-1)

