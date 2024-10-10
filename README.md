## Usage
1. Create a [Spotify app](https://developer.spotify.com/documentation/web-api/tutorials/getting-started#create-an-app). Note the ID and secret provided in your application dashboard.
2. Copy `config.toml.example` to `config.toml`. Edit `config.toml` and set your ID and secret. Make sure that they are enclosed by quotation marks, TOML is strict about denoting that these values are strings.
3. Run the code. This can be done from source by running `go run .` in the root of the repository or by building a binary with `go build .`. The binary accepts one argument, `-dest`, which is the name of the CSV file your liked tracks will be written to.
4. The application will print out a URL of the form `https://accounts.spotify.com/authorize?access_type=offline&client_id=CLIENT_ID&code_challenge=CODE_CHALLENGE&code_challenge_method=S256&redirect_uri=https://yourappurl.com/callback/&response_type=code&scope=user-library-read&state=state`.
5. Copy this URL and open it in a browser. You will be asked to authenticate with Spotify and allow the application access to your account. The only scope requested is `user-library-read` so it cannot make any changes to your account or library.
6. After allowing the application access, you'll be redirected to a new page. The URL will look like `https://likedumper.ndumas.com/callback?code=AUTH_CODE&state=state`. Copy the value of the `code` URL parameter, everything after the `=` and before `&state`.
7. Paste this code into the terminal.
8. The application will iterate through your Liked tracks and write them to a CSV file, either `tracks.csv` or the filename you provided with the `-dest` flag.
