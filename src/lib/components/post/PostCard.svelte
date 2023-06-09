<script lang="ts">
  import type { PostRecord } from '$lib/server/pb.d';
  import { formatDate } from '$lib/dateUtil';
  export let post: PostRecord;
</script>


<div class="card w-full mb-4">
  <header class="card-header">
    <h2 class="sm:text-3xl text-xl"><a href="/blog/{post.slug}">{post.title}</a></h2>
    <div class="text-secondary-600 dark:text-secondary-100">{formatDate(post.created)}</div>
    <div class="flex flex-row">
      {#each (post.expand.tags ? post.expand.tags : []) as tag}
        <div class="card px-1 mr-2 variant-filled-primary"><a href="/tag/{tag.name}">{tag.name}</a></div>
      {/each}
    </div>
  </header>
  <section class="p-4">{post.subTitle}</section>
  <footer class="card-footer text-tertiary-600 dark:text-tertiary-100"><a href="/blog/{post.slug}">Read More...</a></footer>
</div>