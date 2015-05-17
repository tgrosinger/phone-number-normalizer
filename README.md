# Phone Number Normalizer

## Usage

```
sudo apt-get install xsel xclip
chmod +x normalizer
./normalizer
```

With the application running any phone numbers copied to the clipboard will be
converted into `+12345678900` format.

## Example

I made this script then I was converting all my contacts to this format. Here
was the procedure:

1. Open a contact
2. Double click and drag to select the entire number
3. Press `ctrl + v`
4. Go back and open another contact

## Known Issues

To save a step, this script uses `xsel` rather than `xclip` to read input. For
some reason text seems to only be recognized if selected using a double-click
and drag method.
