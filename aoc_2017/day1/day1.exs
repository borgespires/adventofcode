defmodule Main do
    # returns a list of all digits that match the next one
    def match_next_digit(captcha)
    def match_next_digit([]), do: []
    def match_next_digit([h|tail]),
        do: match_next_digit([h|tail] ++ [h], [])
    def match_next_digit([_], acc), do: acc    
    def match_next_digit([d, d|tail], acc),
        do: match_next_digit([d|tail], [d|acc])
    def match_next_digit([_|tail], acc),
        do: match_next_digit(tail, acc)

    # returns a list of all digits that match the digit halfway around the circular list
    def match_halfway_digit(captcha) do
        half = round(length(captcha)/2)
        Enum.split(captcha, half)
            |> match_halfway_digit([])
    end
    def match_halfway_digit({ [], [] }, acc), do: acc
    def match_halfway_digit({ [h|f_tail], [h|s_tail] }, acc), 
        do: match_halfway_digit({ f_tail, s_tail },[h|acc])
    def match_halfway_digit({ [_|f_tail], [_|s_tail] }, acc),
        do: match_halfway_digit({ f_tail, s_tail }, acc)

    def sum_of(captcha, match_fn) do
        captcha
            |> match_fn.()
            |> Enum.sum
    end
end

input = File.read!("input.txt")
    |> String.codepoints
    |> Enum.map(&String.to_integer/1)

#  find the sum of all digits that match the next digit in the list    
input
    |> Main.sum_of(&Main.match_next_digit/1)
    |> IO.puts

# consider the digit halfway around the circular list
input
    |> Main.sum_of(&Main.match_halfway_digit/1)
    |> (&(&1 * 2)).()
    |> IO.puts