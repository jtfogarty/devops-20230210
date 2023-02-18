<script lang="ts">
	import '@skeletonlabs/skeleton/themes/theme-crimson.css';
	import '@skeletonlabs/skeleton/styles/all.css';
	import '../app.postcss';
	import { AppShell, AppBar, AppRail, AppRailTile, Modal, modalStore } from '@skeletonlabs/skeleton';
	import type { ModalSettings, ModalComponent } from '@skeletonlabs/skeleton';
	import ModalExampleForm from '$lib/Modal/ModalExampleForm.svelte';
	import Navigation from '$lib/Navigation/Navigation.svelte';
	import { IconManualGearbox, IconApiApp, IconSettings, IconBrandGitlab, IconApple, } from '@tabler/icons-svelte';
	import data from "./data.json"
	import { goto } from '$app/navigation';

	function handleClick() {
		const text = 'Hello Kitty';
		goto(`/contact?text=${encodeURIComponent(text)}`);
	}
	
	let json = {};
	json = data
	function clickHandler(e) {
		console.log("clickHandler", e);
	}

	function triggerAlert(): void {
		console.log('working');
	}

	function ModalGitLabForm(): void{
		const c: ModalComponent = {ref: ModalExampleForm};
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
</script>

<!-- App Shell -->
<AppShell >
	<svelte:fragment slot="header">
		<!-- App Bar -->
		<AppBar>
			<svelte:fragment slot="lead">
				<strong class="text-xl">Shield Explorer</strong>
			</svelte:fragment>
			<svelte:fragment slot="trail">
				<a class="btn btn-sm variant-ghost-surface" href="/about">About</a>
				<a class="btn btn-sm variant-ghost-surface" href="/contact">contact</a>
				<a class="btn btn-sm variant-ghost-surface" href="/blog">Blog</a>
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
				<Navigation tree={json} onClick={clickHandler}/>
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
			<button on:click={() => window.open('https://www.drudgereport.com', '_blank')} class="btn-icon">Drudge</button>
			<a href="/" class="btn-icon unstyled hover:bg-primary-hover-token" >A</a>
		</div>
	</svelte:fragment>

	<!-- Page Route Content -->
	<slot />
</AppShell>

<Modal />
