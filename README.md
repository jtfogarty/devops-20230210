# Shield Explorer

## About

The Shield Explorer has the following features; 
 - Browse all of the modules and READMEs
 - Generates Application Context YAML
 - Subnet API requests
 - Ensures Application Context and Shield Metadata are sync'd


Run `wails dev` 
To access your Go methods, use http://localhost:34115. 
Connect and call your Go code from devtools.<p/>
You can run `wails dev --host 10.10.51.1:5173` to expose

Run `wails build` to create a distribution

## ToDo's

<input type="checkbox" checked> Create a form or modal to allow user to add gitlab secret. This data should be stored in a svelte:store. <p/>
<input type="checkbox" > Use gitlab secret to connect to gitlab.com and read a readme onto main area.<p/>
<input type="checkbox" > Figure out how to create tabs so when a user clicks on a module one tab is displayed for each element in an array. <p/>
<input type="checkbox"> When user clicks on a module, display code in main area code block.<p/>
<input type="checkbox"> Figure out why the build program does not work.
  