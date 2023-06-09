<script lang="ts">
	import { getYear } from '$lib/dateUtil';

  export let data;
</script>

<svelte:head>
	<title>About</title>
</svelte:head>

<div class="card w-full">
  <header class="card-header">
    <h2 class="sm:text-3xl text-xl">Contact</h2>
  </header>
  <section class="p-4">
    <div class="flex-row flex align-middle justify-center">
      <div class="px-2">
        <button type="button" class="btn-icon btn-icon-xl variant-filled"><a href="https://www.github.com/amirk999" target="_blank"><i class="fa-brands fa-github fa-2xl"></i></a></button>
      </div>
      <div class="px-2">
        <button type="button" class="btn-icon btn-icon-xl variant-filled"><a href="https://www.linkedin.com/in/amirkarbasi" target="_blank"><i class="fa-brands fa-linkedin fa-2xl"></i></a></button>
      </div>
      <div class="px-2">
        <button type="button" class="btn-icon btn-icon-xl variant-filled"><a href="mailto:amir@karbasi.dev"><i class="fa-solid fa-envelope fa-2xl"></i></a></button>
      </div>
    </div>
  </section>
</div>

<div class="card w-full mt-5">
  <header class="card-header">
    <h2 class="sm:text-3xl text-xl">Professional Experience</h2>
  </header>
  <section class="p-4">
    <div class="flex flex-col">
      {#each data.work as work, i}
        <div class="text-xs">{getYear(work.from)} - {#if work.current} PRESENT {:else} {getYear(work.to)} {/if}</div>
        <div class="sm:text-2xl"><span class="text-primary-600 dark:text-primary-300"><strong>{work.title}</strong> / </span><span>{work.institutionName}, {work.location}</span></div>
        <div class="text-sm">
        <ul class="list">
          {#each work.expand.responsibilities as responsibility}
            <li>
              <span>&#x2022; </span>
              <span class="flex-auto">{responsibility.description}</span>
            </li>
          {/each}
        </ul>
        </div>
        {#if i < (data.work.length - 1)}
          <hr class="!border-t-2 my-3" />
        {/if}
      {/each}
    </div>
  </section>
</div>

<div class="card w-full mt-5">
  <header class="card-header">
    <h2 class="sm:text-3xl text-xl">Education</h2>
  </header>
  <section class="p-4">
    <div class="flex flex-col">
      {#each data.education as education, i}
        <div class="text-xs">{getYear(education.from)} - {getYear(education.to)}</div>
        <div class="sm:text-2xl"><span class="text-primary-600 dark:text-primary-300"><strong>{education.title}</strong> / </span><span>{education.institutionName}</span></div>
        <div class="text-sm">{@html education.description}</div>
        {#if i < (data.education.length - 1)}
          <hr class="!border-t-2 my-3" />
        {/if}
      {/each}
    </div>
  </section>
</div>