window.BENCHMARK_DATA = {
  "lastUpdate": 1651345883118,
  "repoUrl": "https://github.com/Infinytum/go-mojito-bunrouter",
  "entries": {
    "Benchmark": [
      {
        "commit": {
          "author": {
            "email": "0skillallluck@pm.me",
            "name": "Cedric Lewe",
            "username": "0SkillAllLuck"
          },
          "committer": {
            "email": "0skillallluck@pm.me",
            "name": "Cedric Lewe",
            "username": "0SkillAllLuck"
          },
          "distinct": true,
          "id": "fe5b8ede9a98cbb86401be271c43737baeab44e6",
          "message": "Remove init func",
          "timestamp": "2022-04-30T21:10:44+02:00",
          "tree_id": "2e86515f93dd18b5f8158fa25386e407d786ad9a",
          "url": "https://github.com/Infinytum/go-mojito-bunrouter/commit/fe5b8ede9a98cbb86401be271c43737baeab44e6"
        },
        "date": 1651345882054,
        "tool": "go",
        "benches": [
          {
            "name": "Benchmark_Router_Handler",
            "value": 454605,
            "unit": "ns/op\t   29429 B/op\t     182 allocs/op",
            "extra": "2224 times\n2 procs"
          },
          {
            "name": "Benchmark_Router_Handler_Not_Found",
            "value": 384104,
            "unit": "ns/op\t   29429 B/op\t     179 allocs/op",
            "extra": "3181 times\n2 procs"
          },
          {
            "name": "Benchmark_Router_Handler_With_Middleware",
            "value": 382188,
            "unit": "ns/op\t   29275 B/op\t     178 allocs/op",
            "extra": "2974 times\n2 procs"
          }
        ]
      }
    ]
  }
}