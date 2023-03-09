<script lang="ts">
	import '@skeletonlabs/skeleton/themes/theme-crimson.css';
	import '@skeletonlabs/skeleton/styles/all.css';
	import '../app.postcss';
	import { AppShell, AppBar, AppRail, AppRailTile, Modal, modalStore } from '@skeletonlabs/skeleton';
	import type { ModalSettings, ModalComponent } from '@skeletonlabs/skeleton';
	import GitlabModalForm from '$lib/Modal/GitLabModal.svelte';
	import Navigation from '$lib/Navigation/Navigation.svelte';
	import { IconManualGearbox, IconApiApp, IconSettings, IconBrandGitlab, IconApple, } from '@tabler/icons-svelte';
	import data from "./data.json"
	import {LaunchChrome} from '../../wailsjs/go/main/App.js';
	import { onMount } from "svelte";

	onMount(() => {
			window.addEventListener("NavClick", handleNavClick);
			
			return () => {
				window.removeEventListener("NavClick", handleNavClick);
			};
		});


	function lnchUrl(url: string): void {
		LaunchChrome(url.toString())
	}
	
	function handleNavClick(event: CustomEvent<string>) {
		const data = event.detail;
    console.log("Parent Clicked:", data);
	}
	let json = {};
	json = data
	function clickHandler(e: any) {
		console.log("hello", e);
	}

	function triggerAlert(): void {
		console.log('working');
	}

	function ModalGitLabForm(): void{
		const c: ModalComponent = {ref: GitlabModalForm};
		const d: ModalSettings = {
			type: 'component',
			title: 'Gitlab Token',
			body: 'Please enter the below and click Submit',
			component: c,
			response: (r: any) => {
				if (r) console.log('response:', r);
			},
			meta: {
				foo: 'bar',
				fizz: 'buzz',
				fn: triggerAlert
			}
		};
		modalStore.trigger(d);
	}

	const navs = [
		{
			title: "Drudge",
			url: "https://drudgereport.com"
		},
		{
			title: "YouTube",
			url: "https://youtube.com"
		},
		{
			title: "LevelUp",
			url: "https://levelup.video"
		}
	]
</script>

<!-- App Shell -->
<AppShell >
	<svelte:fragment slot="header">
		<!-- App Bar -->
		<AppBar>
			<svelte:fragment slot="lead">
				<a href="/" class="flex-shrink-0 text-xl font-bold">
					Shield Explorer
				  </a>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				{#each navs as {title, url}}
					<button class="btn btn-sm variant-ghost-surface" on:click={() => lnchUrl(url)}>{title}</button>
				{/each}
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<svelte:fragment slot="sidebarLeft">
		<div class="flex h-full">
			<div style="background-color: #222F44">
				<AppRail >
					<svelte:fragment slot="lead">
						<AppRailTile tag="a">
							<button >
								<IconApiApp  size={36} stroke={1} color=#ccc/>
							</button>
						</AppRailTile>
						<AppRailTile tag="b" >
							<button >
								<IconManualGearbox  size={36} stroke={1} color=#ccc/>
							</button>
						</AppRailTile>							
						<AppRailTile tag="c" >
							<button on:click={ ModalGitLabForm }>
								<IconBrandGitlab  size={36} stroke={1} color=#ccc/>
							</button>
						</AppRailTile>
						<AppRailTile tag="d" >
							<button>
								<IconApple  size={36} stroke={1} color=#ccc/>
							</button>
						</AppRailTile>													
					</svelte:fragment>
					<svelte:fragment slot="trail">
						<AppRailTile tag="e" >
							<button>
								<IconSettings  size={36} stroke={1} color=#ccc/>
							</button>
						</AppRailTile>
					</svelte:fragment>
				</AppRail>
			</div> 
			<div style="background-color: #1B2636" class="pt-10 px-4">
				<Navigation on:NavClick={ handleNavClick } tree={json}  />
			</div>
		</div>
	</svelte:fragment>
	<svelte:fragment slot="pageHeader" >
		<div class="flex justify-end" style="background-color: #1B2636">
			<button>
				<a href="/about">
					<IconApiApp  size={36} stroke={1} color=#ccc />
				</a>
			</button>
			<button on:click={ ModalGitLabForm } class="px-4">
				<IconManualGearbox class="hover:bg-primary-hover-token" size={36} stroke={1} color=#ccc />
			</button>
		</div>
	</svelte:fragment>

	<!-- Page Route Content -->
	<slot />
</AppShell>

<Modal />
