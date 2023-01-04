# schaar2pod

Converts text chapter marks to json formatted ones

Example: Render chaptermarks from [Ultraschall](https://ultraschall.fm/) and convert it to use it with [Castopod](https://castopod.org/)

## Usage: 
`schaar2pod inputfile.txt`

This will generate a new file called `inputfile.txt.json`

## Example:

Input file:
```
00:00:00.000 Anfang
00:11:50.200 Idee 1: 5-4-3-2-1 der Countdown Podcast
00:17:09.278 Idee 2: Große Stimmen lesen unspannende Dinge
00:20:21.369 Idee 3: Dein Mitbewohner
00:24:13.182 Idee 4: Leute, die keine Lust drauf haben
00:26:52.082 Idee 5: Die Schweigeminute
```

Output file:

```json
{
  "version": "1.2.0",
  "chapters": [
    {
      "startTime": 0,
      "title": "Anfang"
    },
    {
      "startTime": 710.2,
      "title": "Idee 1: 5-4-3-2-1 der Countdown Podcast"
    },
    {
      "startTime": 1029.278,
      "title": "Idee 2: Große Stimmen lesen unspannende Dinge"
    },
    {
      "startTime": 1221.369,
      "title": "Idee 3: Dein Mitbewohner"
    },
    {
      "startTime": 1453.182,
      "title": "Idee 4: Leute, die keine Lust drauf haben"
    },
    {
      "startTime": 1612.082,
      "title": "Idee 5: Die Schweigeminute"
    }
  ]
}
```
