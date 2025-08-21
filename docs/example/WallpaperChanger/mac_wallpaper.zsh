#!/bin/zsh
# This script downloads an image from a URL and sets it as the desktop wallpaper on macOS

IMAGE_URL="http://example.com:8080/random"
WORKDIR="$HOME/dailimage"

if [[ ! -d "$WORKDIR" ]]; then
  mkdir -p "$WORKDIR"
  echo "Created directory: $WORKDIR"
fi

# The file name needs to change or mac will keep using the cached image
bgfile="$WORKDIR/background-$(date -Iseconds).png"
rm -v "$WORKDIR"/background-*.png # remove old image

# Download image
curl "$IMAGE_URL" -o "$bgfile" || {
  echo "Failed to download image from $IMAGE_URL"
  exit 1
}
echo "Downloaded $IMAGE_URL to $bgfile"

# Set the desktop wallpaper with AppleScript
osascript <<EOF
tell application "System Events"
	tell every desktop
		set picture to "$bgfile"
	end tell
end tell
EOF
