class GapBuffer:

	def __init__(self, gap_size=10, buffer_size=10):
		self.buffer = ['_' for i in range(buffer_size + 40)]
		self.gap_left = 0
		self.gap_right = self.gap_left + gap_size - 1
		self.buffer_size = buffer_size

	def insert(self, char):
		'''
		Insert a character at the left of the gap.
		In practice, get the target index and insert the character at the index directly.
		'''
		if self.gap_left == self.gap_right:
			self.expand_gap()

		self.buffer[self.gap_left] = char
		self.gap_left += 1

	def delete(self):
		'''
		Delete the character at the left of the gap.
		In practice, get the target index and delete the character at the index directly.
		'''
		if self.gap_left <= 0:
			return
		
		self.gap_left -= 1
		self.buffer[self.gap_left] = '_'
	
	def move_left(self):
		'''
		Move the gap to the left. It's a method to move the cursor to the left.
		In practice, it's inefficient. Move the gap to the specified position directly instead.
		'''
		self.move_gap(self.gap_left - 1)
	
	def move_right(self):
		'''
		Move the gap to the right. It's a method to move the cursor to the left.
		In practice, it's inefficient. Move the gap to the specified position directly instead.
		'''
		self.move_gap(self.gap_left + 1)

	def move_gap(self, to):
		if to < 0:
			return
		if to >= self.buffer_size:
			return
		if to < self.gap_left:
			for index in range(self.gap_left - to):
				self.buffer[self.gap_right - index] = self.buffer[self.gap_left - index - 1]
			to_move = self.gap_left - to
			self.gap_left -= to_move
			self.gap_right -= to_move
		elif to > self.gap_left:
			for index in range(to - self.gap_left):
				self.buffer[self.gap_left + index] = self.buffer[self.gap_right + index + 1]
			to_move = to - self.gap_left
			self.gap_left += to_move
			self.gap_right += to_move
		
		for index in range(self.gap_left, self.gap_right + 1):
			self.buffer[index] = '_'

		if self.gap_right >= self.buffer_size - 1:
			self.expand_buffer()
		

	def expand_gap(self, to_add=10):
		to_paste_back = self.buffer[self.gap_left:self.gap_left + to_add]

		for index in range(to_add):
			self.buffer[self.gap_left + index] = '_'
		
		for index, char in enumerate(to_paste_back):
			self.buffer[self.gap_left + to_add + index] = char
		
		self.expand_buffer(to_add)
		self.gap_right += to_add
	
	def expand_buffer(self, to_add=40):
		self.buffer += ['_' for i in range(to_add)]
		self.buffer_size += to_add


	def print_buffer(self):
		copied_buffer = self.buffer.copy()
		for index in range(self.gap_left, self.gap_right + 1):
			copied_buffer[index] = "⬜"
		print(copied_buffer[:self.buffer_size])

if __name__ == '__main__':
	gb = GapBuffer()
	gb.print_buffer()
	while True:
		try:
			print('commands:')
			print('q: quit')
			print('i: insert')
			print('d: delete')
			print('r: move right')
			print('l: move left')

			command = input('Enter a command: ')
			if command == 'q':
				break
			elif command == 'i':
				char = input('Enter a character: ')
				if len(char) != 1:
					print('Invalid character.')
					continue
				gb.insert(char)
			elif command == 'd':
				gb.delete()
			elif command == 'r':
				gb.move_right()
			elif command == 'l':
				gb.move_left()
			else:
				print('Invalid command.')
			gb.print_buffer()
			
		except Exception as e:
			print('An error occurred:', e)
			continue