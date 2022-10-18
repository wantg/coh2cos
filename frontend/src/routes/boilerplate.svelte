<script>
  import { beforeUpdate, onMount, afterUpdate } from 'svelte';
</script>

<svelte:head />

<div>
  <div class="p-4">
    <button class="button is-white">White</button>
    <button class="button is-light">Light</button>
    <button class="button is-dark">Dark</button>
    <button class="button is-black">Black</button>
    <button class="button is-text">Text</button>
    <button class="button is-ghost">Ghost</button>
  </div>

  <div class="p-4">
    <div class="container">
      <div class="notification is-primary">
        This container is <strong>centered</strong> on desktop and larger viewports.
      </div>
    </div>

    <div class="container is-widescreen">
      <div class="notification is-primary">
        This container is <strong>fullwidth</strong> <em>until</em> the <code>$widescreen</code> breakpoint.
      </div>
    </div>

    <div class="container is-fullhd">
      <div class="notification is-primary">
        This container is <strong>fullwidth</strong> <em>until</em> the <code>$fullhd</code> breakpoint.
      </div>
    </div>

    <div class="container is-max-desktop">
      <div class="notification is-primary">
        This container has a <code>max-width</code> of <code>$desktop - $container-offset</code> on widescreen and fullhd.
      </div>
    </div>

    <div class="container is-max-widescreen">
      <div class="notification is-primary">
        This container has a <code>max-width</code> of <code>$widescreen - $container-offset</code> on fullhd.
      </div>
    </div>

    <div class="container is-fluid">
      <div class="notification is-primary">
        This container is <strong>fluid</strong>: it will have a 32px gap on either side, on any viewport size.
      </div>
    </div>
  </div>

  <div class="flex-example">
    <p class="title is-5 m-0 pt-5 pb-2">is-flex-direction-row</p>
    <div class="is-flex is-flex-direction-row">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-flex-direction-row-reverse</p>
    <div class="is-flex is-flex-direction-row-reverse">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-flex-direction-column</p>
    <div class="is-flex is-flex-direction-column">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-flex-direction-column-reverse</p>
    <div class="is-flex is-flex-direction-column-reverse">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-flex-wrap-nowrap</p>
    <div class="is-flex is-flex-wrap-nowrap">
      <div><button class="button is-info">nowrap 1</button></div>
      <div><button class="button is-info">nowrap 2</button></div>
      <div><button class="button is-info">nowrap 3</button></div>
      <div><button class="button is-info">nowrap 4</button></div>
      <div><button class="button is-info">nowrap 5</button></div>
      <div><button class="button is-info">nowrap 6</button></div>
      <div><button class="button is-info">nowrap 7</button></div>
      <div><button class="button is-info">nowrap 8</button></div>
      <div><button class="button is-info">nowrap 9</button></div>
      <div><button class="button is-info">nowrap 10</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-flex-wrap-wrap</p>
    <div class="is-flex is-flex-wrap-wrap">
      <div><button class="button is-info">wrap 1</button></div>
      <div><button class="button is-info">wrap 2</button></div>
      <div><button class="button is-info">wrap 3</button></div>
      <div><button class="button is-info">wrap 4</button></div>
      <div><button class="button is-info">wrap 5</button></div>
      <div><button class="button is-info">wrap 6</button></div>
      <div><button class="button is-info">wrap 7</button></div>
      <div><button class="button is-info">wrap 8</button></div>
      <div><button class="button is-info">wrap 9</button></div>
      <div><button class="button is-info">wrap 10</button></div>
      <div><button class="button is-info">wrap 11</button></div>
      <div><button class="button is-info">wrap 12</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-flex-wrap-wrap-reverse</p>
    <div class="is-flex is-flex-wrap-wrap-reverse">
      <div><button class="button is-info">wrap 1</button></div>
      <div><button class="button is-info">wrap 2</button></div>
      <div><button class="button is-info">wrap 3</button></div>
      <div><button class="button is-info">wrap 4</button></div>
      <div><button class="button is-info">wrap 5</button></div>
      <div><button class="button is-info">wrap 6</button></div>
      <div><button class="button is-info">wrap 7</button></div>
      <div><button class="button is-info">wrap 8</button></div>
      <div><button class="button is-info">wrap 9</button></div>
      <div><button class="button is-info">wrap 10</button></div>
      <div><button class="button is-info">wrap 11</button></div>
      <div><button class="button is-info">wrap 12</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-justify-content-flex-start</p>
    <div class="is-flex is-justify-content-flex-start">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-justify-content-flex-end</p>
    <div class="is-flex is-justify-content-flex-end">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-justify-content-center</p>
    <div class="is-flex is-justify-content-center">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-justify-content-space-between</p>
    <div class="is-flex is-justify-content-space-between">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-justify-content-space-around</p>
    <div class="is-flex is-justify-content-space-around">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-justify-content-space-evenly</p>
    <div class="is-flex is-justify-content-space-evenly">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-justify-content-start (Chrome No support)</p>
    <div class="is-flex is-justify-content-start">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-justify-content-end (Chrome No support)</p>
    <div class="is-flex is-justify-content-end">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-justify-content-left (Chrome No support)</p>
    <div class="is-flex is-justify-content-left">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-justify-content-right (Chrome No support)</p>
    <div class="is-flex is-justify-content-right">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-content-flex-start</p>
    <div class="is-flex is-align-content-flex-start">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-content-flex-end</p>
    <div class="is-flex is-align-content-flex-end">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-content-center</p>
    <div class="is-flex is-align-content-center">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-content-space-between</p>
    <div class="is-flex is-align-content-space-between">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-content-space-around</p>
    <div class="is-flex is-align-content-space-around">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-content-space-evenly (Chrome No support)</p>
    <div class="is-flex is-align-content-space-evenly">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-content-stretch</p>
    <div class="is-flex is-align-content-stretch">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-content-start (Chrome No support)</p>
    <div class="is-flex is-align-content-start">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-content-end (Chrome No support)</p>
    <div class="is-flex is-align-content-end">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-content-baseline</p>
    <div class="is-flex is-align-content-baseline">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-items-stretch</p>
    <div class="is-flex is-align-items-stretch">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-items-flex-start</p>
    <div class="is-flex is-align-items-flex-start">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-items-flex-end</p>
    <div class="is-flex is-align-items-flex-end">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-items-center</p>
    <div class="is-flex is-align-items-center">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-items-baseline</p>
    <div class="is-flex is-align-items-baseline">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-items-start</p>
    <div class="is-flex is-align-items-start">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-items-end</p>
    <div class="is-flex is-align-items-end">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-items-self-start</p>
    <div class="is-flex is-align-items-self-start">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-items-self-end</p>
    <div class="is-flex is-align-items-self-end">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-self-auto</p>
    <div class="is-flex is-align-self-auto">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-self-flex-start</p>
    <div class="is-flex is-align-self-flex-start">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-self-flex-end</p>
    <div class="is-flex is-align-self-flex-end">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-self-center</p>
    <div class="is-flex is-align-self-center">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-self-baseline</p>
    <div class="is-flex is-align-self-baseline">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-align-self-stretch</p>
    <div class="is-flex is-align-self-stretch">
      <div><button class="button is-info">1</button></div>
      <div><button class="button is-info">2</button></div>
      <div><button class="button is-info">3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-flex-grow</p>
    <div class="is-flex">
      <div class="is-flex-grow-0"><button class="button is-info">is-flex-grow-0</button></div>
      <div class="is-flex-grow-1"><button class="button is-info">is-flex-grow-1</button></div>
      <div class="is-flex-grow-2"><button class="button is-info">is-flex-grow-2</button></div>
      <div class="is-flex-grow-3"><button class="button is-info">is-flex-grow-3</button></div>
    </div>
    <p class="title is-5 m-0 pt-5 pb-2">is-flex-shrink</p>
    <div class="is-flex">
      <div class="is-flex-shrink-0" style="flex-basis: 400px;"><button class="button is-info">0</button></div>
      <div class="is-flex-shrink-1" style="flex-basis: 400px;"><button class="button is-info">1</button></div>
      <div class="is-flex-shrink-2" style="flex-basis: 400px;"><button class="button is-info">2</button></div>
      <div class="is-flex-shrink-3" style="flex-basis: 400px;"><button class="button is-info">3</button></div>
    </div>
  </div>
</div>

<style>
  .flex-example {
    max-width: 50%;
    margin: 0 auto;
  }
  .is-flex {
    min-height: 100px;
    border: 1px #bbb solid;
  }
  .is-flex > div {
    height: 60px;
    padding: 10px;
  }
</style>
