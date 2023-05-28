const std = @import("std");
const expect = std.testing.expect;
const expectEqual = std.testing.expectEqual;
const Allocator = std.testing.Allocator;

pub fn Stack(comptime T: type) type {
    return struct {
        const Node = struct {
            value: T,
            next: ?*Node,
        };
        head: ?*Node,
        length: usize,
        allocator: std.mem.Allocator,
        const Self = @This();

        pub fn init(allocator: std.mem.Allocator) Self {
            return .{
                .head = null,
                .length = 0,
                .allocator = allocator,
            };
        }

        pub fn push(self: *Self, value: T) !void {
            self.length += 1;
            var newNode = try self.allocator.create(Node);
            newNode.value = value;
            newNode.next = self.head;
            self.head = newNode;
        }

        pub fn pop(self: *Self) ?T {
            if (self.head) |head| {
                self.length -= 1;
                self.head = head.next;
                defer self.allocator.destroy(head);
                return head.value;
            }
            return null;
        }

        pub fn peek(self: Self) ?T {
            if (self.head) |head| {
                return head.value;
            }
            return null;
        }

        pub fn print(self: Self) void {
            std.debug.print("Stack {{", .{});
            var curr = self.head;
            while (curr) |node| : (curr = node.next) {
                std.debug.print("{}", .{node.value});
            }
            std.debug.print("}}\n", .{});
        }
    };
}

pub fn main() !void {
    var alloc = std.heap.GeneralPurposeAllocator(.{}){};
    var allocator = alloc.allocator();
    var stack = Stack(i32).init(allocator);
    defer _ = alloc.deinit();

    try stack.push(1);
    stack.print();
    try stack.push(2);
    stack.print();
    try stack.push(3);
    stack.print();
    _ = stack.pop();
    stack.print();
    _ = stack.pop();
    stack.print();
    _ = stack.pop();
    stack.print();
}

test "Stack test" {
    var stack = Stack(i32).init(std.testing.allocator);

    // Test empty stack
    try expect(stack.length == 0);
    try expectEqual(stack.peek(), null);
    try expectEqual(stack.pop(), null);
    stack.print();

    // Test push and peek
    try stack.push(1);
    try expect(stack.length == 1);
    try expectEqual(stack.peek(), 1);
    stack.print();

    // Test pop
    try expectEqual(stack.pop(), 1);
    try expect(stack.length == 0);
    stack.print();
}
