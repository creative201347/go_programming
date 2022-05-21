import { useState } from "react";
import { useForm } from "@mantine/hooks";
import { Button, Group, Modal, Textarea, TextInput } from "@mantine/core";
import { ENDPOINT, Todo } from "../App";
import { KeyedMutator } from "swr";

const AddTodo = ({ mutate }: { mutate: KeyedMutator<Todo[]> }) => {
  const [open, setOpen] = useState(false);
  const form = useForm({
    initialValues: {
      title: "",
      body: "",
    },
  });

  async function createTodo(values: { title: string; body: string }) {
    const updated = await fetch(`${ENDPOINT}/api/todos`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(values),
    }).then((r) => r.json());

    mutate(updated);
    form.reset();
    setOpen(false);
  }
  return (
    <>
      <Modal opened={open} onClose={() => setOpen(false)} title="Create Todo">
        <form onSubmit={form.onSubmit(createTodo)}>
          <TextInput
            required
            mb={12}
            label="Name"
            placeholder="You're name..."
            {...form.getInputProps("title")}
          />
          <Textarea
            required
            mb={12}
            label="Technology"
            placeholder="Technology you're intrested in..."
            {...form.getInputProps("body")}
          />
          <Button type="submit">Register</Button>
        </form>
      </Modal>

      <Group position="center">
        <Button fullWidth mb={12} onClick={() => setOpen(true)}>
          Get Started
        </Button>
      </Group>
    </>
  );
};

export default AddTodo;
