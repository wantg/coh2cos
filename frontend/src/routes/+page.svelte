<script>
  import { onMount } from 'svelte';
  import constants from '../constants';

  let matchs = [];
  let pickedMatch;
  let leftPlayers = [];
  let rightPlayers = [];
  let tableHeader = ['Drops', 'Total', 'Ratio', 'Losses', 'Wins', 'Streak', 'Level', 'Rank', 'Mode'];

  onMount(async () => {
    runtime.EventsOn('match-updated', (_matchs) => {
      matchs = _matchs;
      if (_matchs.length > 0) {
        loadMatch(_matchs[0].matchID);
      }
    });
    window.go.main.App.UpdateConfig({ lang: navigator.language });
    if (navigator.language.includes('zh')) {
      tableHeader = ['掉线', '总数', '胜率', '输', '赢', '趋势', '级别', '排名', '模式'];
    }
    window.go.main.App.StartLogListener();
  });

  const loadMatch = (matchID) => {
    window.go.main.App.LoadMatchData(matchID)
      .then((result) => {
        // console.log(result);
        pickedMatch = result;
        const _leftPlayers = [];
        const _rightPlayers = [];
        for (const player of pickedMatch.players) {
          if (player.summary) {
            player.summary = JSON.parse(player.summary);
          } else {
            player.summary = { country: '' };
          }
          if (!player.summary.avatar) {
            player.summary.avatar = '/images/coh2-logo.png';
          }
          if (player.stats) {
            player.stats = convertStatData(JSON.parse(player.stats), player.faction);
          } else {
            player.stats = convertStatData({}, player.faction);
          }
          if (player.side == 0) {
            _leftPlayers.push(player);
          } else {
            _rightPlayers.push(player);
          }
        }
        leftPlayers = _leftPlayers;
        rightPlayers = _rightPlayers;
      })
      .catch((err) => {
        console.error(err);
      });
  };

  const convertStatData = (statOrigianl, faction) => {
    const unknownStat = { drops: '-', total: '-', ratio: '-', losses: '-', wins: '-', streak: '-', level: '-', rank: '-', mode: '-' };
    faction = faction.replace(/_/g, '-');
    const data = constants.LEADERBOARD_ID()[faction];
    for (const mode in data) {
      let stat = unknownStat;
      if (statOrigianl[data[mode]]) {
        stat = statOrigianl[data[mode]];
        stat.rank = stat.rank > 0 ? stat.rank : '-';
        stat.level = stat.ranklevel > 0 ? stat.ranklevel : '-';
        stat.total = stat.wins + stat.losses;
        stat.ratio = parseFloat(((stat.wins / stat.total) * 100).toFixed(1)) + '%';
      }
      data[mode] = stat;
    }
    return data;
  };
</script>

<svelte:head>
  <link rel="stylesheet" href="/stylesheets/flag.css" />
  <!--  --><style>
    body {
      user-select: none;
    }
  </style><!--  -->
</svelte:head>

<!-- svelte-ignore a11y-missing-attribute -->
<!-- svelte-ignore security-anchor-rel-noreferrer -->
<!-- svelte-ignore a11y-click-events-have-key-events -->
<div class="is-flex">
  <div class="p-3">
    <nav class="match-triggers panel">
      {#each matchs as match}
        <a class="panel-block is-block has-text-centered" class:is-active={pickedMatch && pickedMatch.matchID == match.matchID} on:click={loadMatch(match.matchID)}>
          {new Date(match.startedAt).toLocaleString()}
        </a>
      {/each}
    </nav>
  </div>
  <div class="p-3 is-flex-grow-1">
    {#if pickedMatch}
      <!-- <h1>{new Date(pickedMatch.startedAt).toLocaleString()}</h1> -->
      <!-- <pre>{JSON.stringify(pickedMatch.players, '', '  ')}</pre> -->
      <div class="is-flex">
        <div class="is-flex-grow-1 pr-2">
          {#each leftPlayers as player, i}
            <div class="player-summary is-flex left is-flex-direction-row-reverse">
              <div><img src="/images/faction-{player.faction.replace(/_/g, '-')}.png" /></div>
              <div class="avatar pl-2 pr-2"><img src={player.summary.avatar} /></div>
              <div>
                <p>{player.alias}</p>
                <p class="f32 has-text-right" title={player.summary.country.toUpperCase()}><span class="flag {player.summary.country}" /></p>
              </div>
            </div>
            <table class="table is-bordered is-striped is-narrow1 is-hoverable is-fullwidth left-player">
              <thead>
                <tr>
                  {#each tableHeader as h}
                    <th>{h}</th>
                  {/each}
                </tr>
              </thead>
              <tbody>
                {#each Object.entries(player.stats) as [mode, stats], i}
                  <tr class={mode}>
                    <td>{stats.drops}</td><td>{stats.total}</td><td>{stats.ratio}</td><td>{stats.losses}</td>
                    <td>{stats.wins}</td>
                    <td>{@html stats.streak > -1 ? '&nbsp' + stats.streak : stats.streak}</td>
                    <td>{stats.level}</td><td>{stats.rank}</td><td>{mode}</td>
                  </tr>
                {/each}
              </tbody>
            </table>
          {/each}
        </div>
        <div class="is-flex-grow-1 pl-2">
          {#each rightPlayers as player, i}
            <div class="player-summary is-flex right">
              <div><img src="/images/faction-{player.faction.replace(/_/g, '-')}.png" /></div>
              <div class="avatar pl-2 pr-2"><img src={player.summary.avatar} /></div>
              <div>
                <p>{player.alias}</p>
                <p class="f32" title={player.summary.country.toUpperCase()}><span class="flag {player.summary.country}" /></p>
              </div>
            </div>
            <table class="table is-bordered is-striped is-narrow1 is-hoverable is-fullwidth right-player">
              <thead>
                <tr>
                  {#each JSON.parse(JSON.stringify(tableHeader)).reverse() as h}
                    <th>{h}</th>
                  {/each}
                </tr>
              </thead>
              <tbody>
                {#each Object.entries(player.stats) as [mode, stats], i}
                  <tr class={mode}>
                    <td>{mode}</td><td>{stats.rank}</td><td>{stats.level}</td>
                    <td>{@html stats.streak > -1 ? '&nbsp' + stats.streak : stats.streak}</td>
                    <td>{stats.wins}</td><td>{stats.losses}</td><td>{stats.ratio}</td><td>{stats.total}</td><td>{stats.drops}</td>
                  </tr>
                {/each}
              </tbody>
            </table>
          {/each}
        </div>
      </div>
    {/if}
  </div>
</div>

<style>
  .match-triggers {
    width: 180px !important;
    max-height: 100vh;
    overflow-y: auto;
  }
  .panel-block:first-child {
    border-top-left-radius: 6px;
    border-top-right-radius: 6px;
  }
  .panel-block.is-active {
    background-color: #3082c5;
    color: #fff;
  }
  .player-summary .avatar > img {
    min-width: 64px;
    min-height: 64px;
    width: 64px;
    height: 64px;
    border: 1px solid #dbdbdb;
    border-radius: 4px;
    padding: 2px;
    display: block;
  }
  th,
  td {
    text-align: center;
  }
</style>
