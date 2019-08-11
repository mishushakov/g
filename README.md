# g: The Google Assistant command

![](https://i.imgur.com/tHhGezw.gif)

Run queries against Google Assistant on every shell 🎉

This is one of my weekend-research projects. It uses the official [Google Assistant Embedded gRPC API](https://developers.google.com/assistant/sdk/reference/rpc/google.assistant.embedded.v1alpha2) and its [Go Bindings](https://godoc.org/google.golang.org/genproto/googleapis/assistant/embedded/v1alpha2)

Note: the software might break and there is no guarantee it will even work, because the APIs its based on are still **in alpha**. Use with caution

With some little changes to the code-base, you will also be able to output HTML:

![](https://i.imgur.com/BqmAZnD.png)
![](https://i.imgur.com/Tau7mqf.jpg)
![](https://i.imgur.com/7dX2vKT.jpg)
![](https://i.imgur.com/IES145i.jpg)

Ps. The code is not intended for the commercial use.

## Installation

- Visit [Actions on Google](https://console.actions.google.com)

![](https://i.imgur.com/9cgC73I.png)

- Press on "New project" and select your Google Cloud Project (or make a new one) and commit with "Create project"

![](https://i.imgur.com/IFKW2Eh.png)

- Scroll down and press on "Device registration"

![](https://i.imgur.com/CRYc8B7.png)

- Then, press on "REGISTER MODEL"

![](https://i.imgur.com/NrksmWu.png)

- Enter the required info and press "REGISTER MODEL"

![](https://i.imgur.com/KLISzKn.png)


- Press on "Download OAuth 2.0 credentials"

![](https://i.imgur.com/atBhGFn.png)

- You can skip the device traits part
- Make sure you have [Python](https://www.python.org) installed
- Install the google-auth-oauthlib tool `pip install --upgrade "google-auth-oauthlib[tool]"`
- Run the tool

```sh
google-oauthlib-tool --client-secrets <PATH TO YOUR OAUTH2 CREDENTIALS> \
                     --credentials credentials.json \
                     --scope https://www.googleapis.com/auth/assistant-sdk-prototype \
                     --save
```

- Sign in with Google and grant the rights, and you should see the file `credentials.json` where you ran the command
- Download `g` for your OS from the [Releases Page](https://github.com/mishushakov/g/releases)
- Create a new file, called `config.yaml` and populate it with the contents of your device & oauth credentials:

```yaml
oauth: # OAuth Configuration
  client_id: <CLIENT_ID> # OAuth Client ID
  client_secret: <CLIENT_SECRET> # OAuth Client Secret
  scopes: # OAuth scopes
    - https://www.googleapis.com/auth/assistant-sdk-prototype
  auth_url: https://accounts.google.com/o/oauth2/auth # OAuth Auth URL
  token_url: https://accounts.google.com/o/oauth2/token # OAuth Token URL
  refresh_token: <REFRESH_TOKEN> # OAuth Refresh Token
device: # Device Configuration
  endpoint: embeddedassistant.googleapis.com:443 # Service Endpoint
  device_id: default # Device ID
  device_model_id: default # Device Model ID
  language_code: en-US # Language Code
```

- Run the `g <query>` command
- Profit

![](https://i.imgur.com/mbcKtEs.png)

## Development

- Get [Go](https://golang.org/dl/)
- Build with the `build.sh` script